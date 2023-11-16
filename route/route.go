package route

import (
	"capstone-project/config"
	"capstone-project/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(config.JWT_KEY)))

	// Route / to handler function
	// user route
	e.POST("/users/signup", controller.SignUpUserController)
	e.POST("/users/login", controller.LoginUserController)
	e.POST("/users/change-password", controller.ChangeUserPasswordController)

	// need authentication
	// user route
	e.GET("/patients", controller.GetPatientsController)
	e.GET("/patients/:id", controller.GetPatientController)
	e.POST("/patients", controller.CreatePatientController)
	e.PUT("/patients/:id", controller.UpdatePatientController)
	e.DELETE("/patients/:id", controller.DeletePatientController)

	return e
}
