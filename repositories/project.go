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
