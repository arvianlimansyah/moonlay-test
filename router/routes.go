package router

import (
	"github.com/arvianlimansyah/moonlay-test/utils"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type dbConnection struct {
	DB *gorm.DB
}

var (
	globalConn dbConnection
)

func New() *echo.Echo {
	conn := utils.ConnectDB()
	globalConn.DB = conn
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
	app := globalConn

	// Route / to handler function
	e.GET("/health-check", app.HealthCheck)

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
