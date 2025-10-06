package taskRepo

import (
	"portfolioAPI/database"
	journeyModels "portfolioAPI/journey/models"

	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{
		db: database.GetDBClient(),
	}
}

func (repo *TaskRepo) Insert(task journeyModels.Task) error {
	result := repo.db.Create(&task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *TaskRepo) DeleteTasks(journeyId uint) error {
	result := repo.db.Where("experience_id = ?", journeyId).Delete(&journeyModels.Task{})

	return result.Error
}
