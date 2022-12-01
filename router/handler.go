package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arvianlimansyah/moonlay-test/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type dbConnection struct {
	DB *gorm.DB
}

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	resp := models.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}

func (dbc *dbConnection) GetTask(c echo.Context) error {
	var tasks []models.Task
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")

	query := "select t.* from task t order by t.id desc offset " + page + " limit " + limit
	err := dbc.DB.Raw(query).Scan(&tasks).Error
	if err != nil {
		log.Fatalf("Record not found %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	return c.JSON(http.StatusOK, tasks)
}

func (dbc *dbConnection) GetAllTask(c echo.Context) error {
	var tasks []models.Task
	dbc.DB.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}

func (dbc *dbConnection) GetTaskWithSubtask(c echo.Context) error {
	var resp models.TaskWithSubtask
	var task models.Task
	var subtasks []models.Subtask
	taskId := c.QueryParam("id")

	queryTask := "select t.* from task t where t.id = " + taskId
	err := dbc.DB.Raw(queryTask).Scan(&task).Error
	if err != nil {
		log.Fatalf("Record not found %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	querySubtask := "select s.* from subtask s where s.task_id = " + taskId
	err = dbc.DB.Raw(querySubtask).Scan(&subtasks).Error
	if err != nil {
		log.Fatalf("Record not found %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	resp.Title = task.Title
	resp.Description = task.Description
	resp.Attachment = task.Attachment
	resp.Subtasks = subtasks

	return c.JSON(http.StatusOK, resp)
}

func (dbc *dbConnection) CreateTask(c echo.Context) error {
	var task models.Task
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&task)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	dbc.DB.Create(&task)
	return c.JSON(http.StatusOK, task)
}

func (dbc *dbConnection) UpdateTask(c echo.Context) error {
	var input models.Task
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	var task models.Task
	if err := dbc.DB.Where("id = ?", input.ID).First(&task).Error; err != nil {
		log.Fatalf("Record not found %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	dbc.DB.Model(&task).Updates(input)
	return c.JSON(http.StatusOK, task)
}

func (dbc *dbConnection) DeleteTask(c echo.Context) error {
	id := c.QueryParam("id")

	var task models.Task
	if err := dbc.DB.Where("id = ?", id).First(&task).Error; err != nil {
		log.Fatalf("Record not found %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	dbc.DB.Delete(&task)
	return c.JSON(http.StatusOK, task)
}

func (dbc *dbConnection) GetSubtask(c echo.Context) error {
	taskId := c.QueryParam("task_id")
	var subtasks []models.Subtask

	query := "select s.* from subtask s where s.task_id = " + taskId
	err := dbc.DB.Raw(query).Scan(&subtasks).Error
	if err != nil {
		log.Fatalf("Record not found %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	return c.JSON(http.StatusOK, subtasks)
}

func (dbc *dbConnection) CreateSubtask(c echo.Context) error {
	var subtask models.Subtask
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&subtask)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	dbc.DB.Create(&subtask)
	return c.JSON(http.StatusOK, subtask)
}

func (dbc *dbConnection) UpdateSubtask(c echo.Context) error {
	var input models.Subtask
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	var subtask models.Subtask
	if err := dbc.DB.Where("id = ?", input.ID).First(&subtask).Error; err != nil {
		log.Fatalf("Record not found %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	dbc.DB.Model(&subtask).Updates(input)
	return c.JSON(http.StatusOK, subtask)
}

func (dbc *dbConnection) DeleteSubtask(c echo.Context) error {
	id := c.QueryParam("id")

	var subtask models.Subtask
	if err := dbc.DB.Where("id = ?", id).First(&subtask).Error; err != nil {
		log.Fatalf("Record not found %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	dbc.DB.Delete(&subtask)
	return c.JSON(http.StatusOK, subtask)
}
