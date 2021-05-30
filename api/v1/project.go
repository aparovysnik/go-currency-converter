package v1

import (
	"github.com/aparovysnik/go-currency-converter/models"
	"github.com/aparovysnik/go-currency-converter/repositories"
	"github.com/gin-gonic/gin"
)

type Project interface {
	Register(*gin.Context) error
}

type project struct {
	repository repositories.Project
}

func InitProjectController(ginEngine *gin.Engine, repository repositories.Project) {
	project := project{
		repository: repository,
	}
	ginEngine.POST("/project", project.Register)
}

func (project *project) Register(c *gin.Context) {
	err := project.repository.Add(models.Project{})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong.",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Value added.",
		})
	}
}
