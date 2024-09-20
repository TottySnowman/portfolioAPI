package project_repo

import (
	"portfolioAPI/database"
	projectModel "portfolioAPI/project/models"

	"gorm.io/gorm"
)

type Project_Repo struct {
	db *gorm.DB
}

func NewProjectRepo() *Project_Repo {
	return &Project_Repo{
		db: database.GetDBClient(),
	}
}

func (repo *Project_Repo) GetAllProjects() []projectModel.Project {
	var projects []projectModel.Project

	result := repo.db.Find(&projects)
	if result.Error != nil {
    return nil
	}
	return projects
}
