package main

import (
	"github.com/wscherfel/fitlogic-backend/access"
	"github.com/wscherfel/fitlogic-backend/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/wscherfel/fitlogic-backend/controllers"
	"github.com/spf13/viper"
)

func main() {
	db, err := access.ConnectToDb()
	if err != nil {
		panic(err)
	}

	// only for testing purposes, clean DB at every start
	// db.DropTableIfExists(&models.User{}, &models.Project{}, &models.Risk{}, &models.CounterMeasure{})

	// migrate the DB models
	db.AutoMigrate(&models.User{}, &models.Project{}, &models.Risk{}/*, &models.CounterMeasure{}*/)

	viper.SetConfigName("fitlogic-conf")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	// create data access objects
	userDao := access.NewUserDAO(db)
	projectDao := access.NewProjectDAO(db)
	riskDao := access.NewRiskDAO(db)
	//cmDao := access.NewCounterMeasureDAO(db)

	// create controllers
	userController := controllers.NewUserController(
		controllers.UserControllerConfig{
			UserDao: userDao,
		})

	projectController := controllers.NewProjectController(
		controllers.ProjectControllerConfig{
			ProjectDao: projectDao,
			UserDao: userDao,
			RiskDao: riskDao,
		})

	riskController := controllers.NewRiskController(
		controllers.RiskControllerConfig{
			RiskDao: riskDao,
			//CmDao: cmDao,
			ProjectDao: projectDao,
			UserDao: userDao,
		},
	)

	/*cmControlelr := controllers.NewCounterMeasureController(
		controllers.CmControllerConfig{
			CmDao: cmDao,
		},
	)*/

	// currently only endpoint that does not use JWT authentication
	e.POST("/login", userController.Login)

	secret := []byte(viper.GetString("Secret"))
	// route user endpoints
	users := e.Group("/users", middleware.JWT(secret))

	users.POST("/", userController.Create)
	users.GET("/", userController.Read)
	users.GET("/:id", userController.ReadByID)
	users.DELETE("/:id", userController.DeleteByID)
	users.PUT("/:id", userController.UpdateByID)
	users.POST("/:id/changepassword", userController.ChangePasswordByID)

	// route project endpoints
	projects := e.Group("/projects", middleware.JWT(secret))

	projects.POST("/", projectController.Create)
	projects.GET("/", projectController.GetAll)
	projects.POST("/:id/assignusers", projectController.AssignUsers)
	projects.POST("/:id/unassignusers", projectController.UnAssignUsers)
	projects.POST("/:id/assignrisks", projectController.AssignRisks)
	projects.POST("/:id/unassignrisks", projectController.UnAssignRisks)
	projects.GET("/:id", projectController.ReadByID)
	projects.PUT("/:id", projectController.UpdateByID)
	projects.DELETE("/:id", projectController.DeleteByID)
	projects.POST("/risks", projectController.GetRisksOfProjects)

	// route risk endpoints
	risks := e.Group("/risks", middleware.JWT(secret))

	risks.POST("/", riskController.Create)
	risks.GET("/", riskController.GetAll)
	risks.GET("/:id", riskController.ReadByID)
	risks.PUT("/:id", riskController.UpdateByID)
	risks.DELETE("/:id", riskController.DeleteByID)
	/*risks.POST("/:id/assigncms", riskController.AssignCms)
	risks.POST("/:id/unassigncms", riskController.UnAssignCms)

	cms := e.Group("/cms", middleware.JWT(secret))

	cms.POST("/", cmControlelr.Create)
	cms.GET("/", cmControlelr.GetAll)
	cms.PUT("/:id", cmControlelr.UpdateByID)
	cms.DELETE("/:id", cmControlelr.DeleteByID)*/

	e.Logger.Fatal(e.Start("0.0.0.0:"+viper.GetString("Port")))
}
