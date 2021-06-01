package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type health struct {
}

func InitHealthController(ginEngine *gin.Engine) {
	health := &health{}
	ginEngine.GET("/health", health.GetStatus)
}

func (health *health) GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Service is running.",
	})
}
