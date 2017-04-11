package main

import (
	"github.com/wscherfel/fitlogic-backend/access"
	"github.com/wscherfel/fitlogic-backend/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/wscherfel/fitlogic-backend/controllers"
)

const (
	serverPort = "8040"

)

func main() {
	db, err := access.ConnectToDb()
	if err != nil {
		panic(err)
	}

	// only for testing purposes, clean DB at every start
	db.DropTableIfExists(&models.User{}, &models.Project{}, &models.Risk{}, &models.CounterMeasure{})

	db.AutoMigrate(&models.User{}, &models.Project{}, &models.Risk{}, &models.CounterMeasure{})

	e := echo.New()
	e.Use(middleware.Logger())

	userDao := access.NewUserDAO(db)
	userController := controllers.NewUserController(
		controllers.UserControllerConfig{
			UserDao: userDao,
		})

	// currently only endpoint that does not use JWT authentication
	e.POST("/login", userController.Login)

	users := e.Group("/users", /*middleware.JWT([]byte(fitlogic.Secret))*/)

	users.POST("/", userController.Register)

	e.Logger.Fatal(e.Start("0.0.0.0:"+serverPort))
}
