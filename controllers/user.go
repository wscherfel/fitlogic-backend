package controllers

import (
	"github.com/wscherfel/fitlogic-backend/access"
	"github.com/labstack/echo"
	"github.com/wscherfel/fitlogic-backend/models"
	"net/http"
	"github.com/wscherfel/fitlogic-backend/common"
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

func (c *UserController) Register(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}
