package services

import (
	"github.com/aparovysnik/go-currency-converter/models"
	"github.com/aparovysnik/go-currency-converter/repositories"
	utils "github.com/aparovysnik/go-currency-converter/utils"
)

type Project interface {
	Register(contactEmail string) (string, error)
}

type project struct {
	repository repositories.Project
}

func NewProjectService(repository repositories.Project) Project {
	return &project{
		repository: repository,
	}
}

func (srv *project) Register(contactEmail string) (string, error) {
	//Generate API key
	apiKey, err := utils.GenerateRandomString(32)
	if err != nil {
		return "", err
	}

	//Compute API key hash
	apiKeyHash := utils.ComputeHash(apiKey)

	newProject := models.Project{
		ContactEmail: contactEmail,
		ApiKeyHash:   apiKeyHash,
	}

	//Store the new project
	err = srv.repository.Add(newProject)

	if err != nil {
		return "", err
	}

	return apiKey, nil
}
