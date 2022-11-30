package router

import (
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// //create groups
	// adminGroup := e.Group("/admin")

	// //set all middlewares
	e.Use(serverHeader)
	// middlewares.SetAdminMiddlewares(adminGroup)

	//set main routes
	MainGroup(e)

	return e
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Custom-Header", "blah!!!")
		return next(c)
	}
}

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("/health-check", HealthCheck)

	// e.GET("/cats/:data", HealthCheck)
	// e.POST("/cats", handlers.AddCat)

}
