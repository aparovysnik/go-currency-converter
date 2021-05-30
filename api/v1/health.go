package v1

import "github.com/gin-gonic/gin"

type health struct {
}

func InitHealthController(ginEngine *gin.Engine) {
	health := &health{}
	ginEngine.GET("/health", health.GetStatus)
}

func (health *health) GetStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Service is running.",
	})
}
