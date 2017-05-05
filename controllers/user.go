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
	// DefaultAdmin is a structure of a default admin
	// password is "qwerty" hashed with md5 (frontend uses this hashing method)
	DefaultAdmin = &models.User{
		Name: "admin",
		Email: "admin@admin.com",
		Password: "d8578edf8458ce06fbc5bb76a58c5ca4",
		Role: models.RoleAdmin,
	}
)

type UserControllerConfig struct {
	UserDao *access.UserDAO
}

// UserController is a controller that handles user endpoints
type UserController struct {
	UserControllerConfig
}	

// LoginCredentials is structure of request when logging in
type LoginCredentials struct {
	Email string `valid:"email"`
	Password string `valid:"required"`
}

// LoginResponse is structure of response to successful login
type LoginResponse struct {
	ID uint
	Token string
	Name string
	Role int
}

// ChangePasswordRequest is a structure of a request to change password
type ChangePasswordRequest struct {
	OldPassword string `valid:"required"`
	NewPassword string `valid:"required"`
}

// UpdateRequest is a structure of request to update user
type UpdateRequest struct {
	Name string
	Email string `valid:"email"`
	Role int
	Skills string
	Status string
}

func NewUserController(config UserControllerConfig) *UserController {
	newController :=  &UserController{
		UserControllerConfig: config,
	}

	// create the default admin
	newController.UserDao.Create(DefaultAdmin)
	return newController
}

// Login will check credentials in request and log user in if credentials are ok
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
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrWrongEmailOrPassword))
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

// Create will create a new user in DB
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
	user.Password = ""

	return ctx.JSON(http.StatusOK, &user)
}

// Read will get all users from DB and return them
func (c *UserController) Read(ctx echo.Context) error {
	_, _, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}

	users, err := c.UserDao.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}
	for i := range users {
		users[i].Password = ""
	}

	return ctx.JSON(http.StatusOK, users)
}

// ReadByID will return user with details (and associations) if logged user has
// sufficient privileges
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
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}
	user.Risks, err = c.UserDao.GetAllAssociatedRisks(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}
	user.Password = ""

	return ctx.JSON(http.StatusOK, user)
}

// DeleteByID will delete user with ID in path if logged user has sufficient
// privileges (is admin in this instance)
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

	// check if only admin is going to be deleted
	others, err := c.UserDao.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}
	onlyAdmin := true
	for i := range others{
		if others[i].ID != pathID && others[i].Role == models.RoleAdmin {
			onlyAdmin = false
		}
	}
	if onlyAdmin {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrCannotDeleteOnlyAdmin))
	}

	err = c.UserDao.Delete(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}

	return ctx.NoContent(http.StatusOK)
}

// UpdateByID will update user with ID in path
// to new values sent in request body
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
		Status: requestValues.Status,
	}

	if JWTRole == models.RoleAdmin {
		updatedVals.Role = requestValues.Role

		oldVals, err := c.UserDao.ReadByID(pathID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
		}
		// updated user is manager or admin and wants to be downgraded to user
		if oldVals.Role <= models.RoleManager && updatedVals.Role > models.RoleManager {
			projects, err := c.UserDao.GetAllAssociatedProjects(oldVals)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
			}
			for i := range projects {
				if projects[i].ManagerID == pathID {
					return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrManagerStillLeadsProjects))
				}
			}
		}
	}

	newVals, err := c.UserDao.Update(updatedVals, pathID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, common.CreateError(err))
	}
	newVals.Password = ""

	return ctx.JSON(http.StatusOK, newVals)
}

// ChangePasswordByID will change user's password if he sends correct old password
func (c *UserController) ChangePasswordByID(ctx echo.Context) error{
	pathIDuint64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(common.ErrIdInPathWrongFormat))
	}
	pathID := uint(pathIDuint64)
	jwtID, _, err := common.GetUserIdAndRoleFromToken(ctx.Get("user").(*jwt.Token))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, common.CreateError(err))
	}
	user, err := c.UserDao.ReadByID(pathID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, common.CreateError(err))
	}

	if user.ID != jwtID {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrUnsufficientPrivileges))
	}

	req := ChangePasswordRequest{}
	err = common.BindAndValid(ctx, &req)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, common.CreateError(err))
	}

	if req.OldPassword != user.Password {
		return ctx.JSON(http.StatusUnauthorized, common.CreateError(common.ErrWrongPassword))
	}

	user.Password = req.NewPassword

	user, err = c.UserDao.Update(user, pathID)

	return ctx.NoContent(http.StatusOK)
}
