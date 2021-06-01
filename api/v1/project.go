package v1

import (
	"net/http"

	"github.com/aparovysnik/go-currency-converter/api/v1/models"
	"github.com/aparovysnik/go-currency-converter/services"
	"github.com/gin-gonic/gin"
)

type Project interface {
	Register(*gin.Context) error
}

type project struct {
	service services.Project
}

func InitProjectController(ginEngine *gin.Engine, service services.Project) {
	project := project{
		service: service,
	}

	ginEngine.POST("/project", project.Register)
}

func (project *project) Register(c *gin.Context) {
	reqBody := new(models.AddProjectRequest)
	err := c.Bind(reqBody)

	if err != nil || !reqBody.IsValid() {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			BaseResponse: models.BaseResponse{
				Status: http.StatusBadRequest,
			},
			Message: "Invalid request",
		})
		return
	}

	apiKey, err := project.service.Register(reqBody.ContactEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			BaseResponse: models.BaseResponse{
				Status: http.StatusInternalServerError,
			},
			Message: "Something went wrong.",
		})
		return
	}

	c.JSON(http.StatusOK, models.AddProjectResponse{
		BaseResponse: models.BaseResponse{
			Status: http.StatusOK,
		},
		ApiKey: apiKey,
	})
}
