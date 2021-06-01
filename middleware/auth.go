package middleware

import (
	"net/http"

	"github.com/aparovysnik/go-currency-converter/api/v1/models"
	"github.com/aparovysnik/go-currency-converter/repositories"
	"github.com/aparovysnik/go-currency-converter/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(projectRepo repositories.Project) gin.HandlerFunc {
	return func(c *gin.Context) {
		success := true
		apiKey := c.Query("apiKey")
		if apiKey == "" {
			success = false
		} else {
			apiKeyHash := utils.ComputeHash(apiKey)

			if !projectRepo.HasProjectWithApiKey(apiKeyHash) {
				success = false
			}
		}

		if !success {
			c.JSON(http.StatusUnauthorized, models.StatusResponse{
				BaseResponse: models.BaseResponse{
					Status: http.StatusUnauthorized,
				},
				Message: "Unauthorized request.",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
