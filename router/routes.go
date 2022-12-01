package router

import (
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	e.Use(serverHeader)
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
	var app dbConnection

	// Route / to handler function
	e.GET("/health-check", HealthCheck)

	e.GET("/task/:data", app.GetTask)
	e.GET("/task-all", app.GetAllTask)
	e.GET("/task-subtask/:data", app.GetTaskWithSubtask)
	e.POST("/task-create", app.CreateTask)
	e.POST("/task-update", app.UpdateTask)
	e.GET("/task-delete/:data", app.DeleteTask)

	e.GET("/subtask/:data", app.GetSubtask)
	e.POST("/subtask-create", app.CreateSubtask)
	e.POST("/subtask-update", app.UpdateSubtask)
	e.GET("/subtask-delete/:data", app.DeleteSubtask)

}
