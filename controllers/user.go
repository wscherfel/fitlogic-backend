package controllers

import (
	"github.com/wscherfel/fitlogic-backend/access"
	"github.com/labstack/echo"
	"github.com/wscherfel/fitlogic-backend/models"
	"net/http"
	"github.com/wscherfel/fitlogic-backend/common"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

var(
	DefaultAdmin = &models.User{
		Name: "admin",
		Email: "admin@admin.com",
		Password: "qwerty", // change this to hashed default password
		Role: models.RoleAdmin,
	}
)

type UserControllerConfig struct {
	UserDao *access.UserDAO
}

type UserController struct {
	UserControllerConfig
}

type LoginCredentials struct {
	Email string `valid:"email"`
	Password string `valid:"required"`
}

type LoginResponse struct {
	ID uint
	Token string
	Name string
	Role int
}

type UpdateRequest struct {
	Name string
	Email string `valid:"email"`
	Role int
	Skills string
}

func NewUserController(config UserControllerConfig) *UserController {
	newController :=  &UserController{
		UserControllerConfig: config,
	}

	// create the default admin
	newController.UserDao.Create(DefaultAdmin)
	return newController
}

func (c *UserController) Login(ctx echo.Context) error {
	credentials := LoginCredentials{}
	err := common.BindAndValid(ctx, &credentials)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	read, err := c.UserDao.ReadByEmail(credentials.Email)
	// error during read from DB
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	// no users with given email found or password does not match
	if len(read) == 0 || read[0].Password != credentials.Password {
		return ctx.JSON(http.StatusUnauthorized, common.ErrWrongEmailOrPassword)
	}

	tokenString, err := common.CreateToken(read[0].ID, read[0].Role)
	// token generation error - weird stuff happened
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	response := &LoginResponse{
		ID: read[0].ID,
		Token: tokenString,
		Name: read[0].Name,
		Role: read[0].Role,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *UserController) Create(ctx echo.Context) error {
	_, role, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if role != models.RoleAdmin {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}
	user := models.User{}
	err = common.BindAndValid(ctx, &user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	user.Projects = []models.Project{}
	user.Risks = []models.Risk{}

	err = c.UserDao.Create(&user)
	// error during create
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, &user)
}

func (c *UserController) Read(ctx echo.Context) error {
	_, role, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if role > models.RoleManager {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	users, err := c.UserDao.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, users)
}

func (c *UserController) ReadByID(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	jwtID, role, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if role >= models.RoleUser && jwtID != pathID{ // the >= condition is for possibility of adding new user roles
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	user, err := c.UserDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	user.Projects, err = c.UserDao.GetAllAssociatedProjects(user)
	/*if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}*/
	user.Risks, err = c.UserDao.GetAllAssociatedRisks(user)
	/*if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}*/

	return ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteByID(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	_, role, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if role != models.RoleAdmin {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	user, err := c.UserDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, common.CreateError(err))
	}

	err = c.UserDao.Delete(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *UserController) UpdateByID(ctx echo.Context) error {
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	jwtID, JWTRole, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	if JWTRole >= models.RoleUser && jwtID != pathID{ // the >= condition is for possibility of adding new user roles
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	requestValues := &UpdateRequest{}
	err = common.BindAndValid(ctx, requestValues)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	updatedVals := &models.User{
		Name: requestValues.Name,
		Email: requestValues.Email,
		Skills: requestValues.Skills,
	}

	if JWTRole == models.RoleAdmin {
		updatedVals.Role = requestValues.Role
	}

	newVals, err := c.UserDao.Update(updatedVals, pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.JSON(http.StatusOK, newVals)
}
