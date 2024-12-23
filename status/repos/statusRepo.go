package statusRepo

import (
	"portfolioAPI/database"
	projectModel "portfolioAPI/project/models"
	statusModels "portfolioAPI/status/models"

	"gorm.io/gorm"
)


type StatusRepo struct {
	db *gorm.DB
}

func NewStatusRepo() *StatusRepo {
	return &StatusRepo{
		db: database.GetDBClient(),
	}
}

func (repo *StatusRepo) GetAllStatus() []statusModels.ProjectStatusDisplay{
  var dbStatus []projectModel.ProjectStatus
  var displayTags []statusModels.ProjectStatusDisplay

  result := repo.db.Find(&dbStatus)

  if result.Error != nil{
    return nil
  }

  for _, status := range dbStatus{
  displayTags = append(displayTags, statusModels.ProjectStatusDisplay{
      StatusID: int(status.ID),
      Status: status.Status,
    })
  }
  
  return displayTags
}
