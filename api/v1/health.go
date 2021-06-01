package v1

import (
	"net/http"

	"github.com/aparovysnik/go-currency-converter/api/v1/models"
	"github.com/gin-gonic/gin"
)

type health struct {
}

func InitHealthController(ginEngine *gin.Engine) {
	health := &health{}
	ginEngine.GET("/health", health.GetStatus)
}

// GetStatus godoc
// @Summary Gets the status of the server.
// @Description get the status indicating if the server is operational.
// @Tags Health
// @Accept */*
// @Produce json
// @Success 200 {object} models.StatusResponse{}
// @Router /health [get]
func (health *health) GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, models.StatusResponse{
		BaseResponse: models.BaseResponse{
			Status: http.StatusOK,
		},
		Message: "Service is running.",
	})
}
