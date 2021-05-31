package repositories

import (
	"github.com/aparovysnik/go-currency-converter/models"
	"gorm.io/gorm"
)

type project struct {
	db *gorm.DB
}

type Project interface {
	Add(project models.Project) error
	HasProjectWithApiKey(apiKeyHash []byte) bool
}

func NewProjectRepository(db *gorm.DB) Project {
	return &project{
		db: db,
	}
}

// Add adds a single project to the collection
func (repo *project) Add(project models.Project) error {
	if err := repo.db.Create(&project).Error; err != nil {
		return err
	}

	return nil
}

// Add adds a single project to the collection
func (repo *project) HasProjectWithApiKey(apiKeyHash []byte) bool {
	var project models.Project
	if err := repo.db.First(&project, "api_key_hash = ?", apiKeyHash).Error; err != nil {
		return false
	}

	return true
}
