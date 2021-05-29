package router

import (
	v1 "github.com/aparovysnik/go-currency-converter/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", v1.GetStatus)

	return r
}
