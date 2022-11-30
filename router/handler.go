package router

import (
	"net/http"

	"github.com/arvianlimansyah/moonlay-test/models"

	"github.com/labstack/echo"
)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	//requestID := uuid.NewV4()
	//c.Logger().Infof("RequestID: %s", requestID)

	resp := models.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
