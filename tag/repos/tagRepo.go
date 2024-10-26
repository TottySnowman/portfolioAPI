package tagRepo

import (
	"portfolioAPI/database"
	tagModel "portfolioAPI/tag/models"

	"gorm.io/gorm"
)

type TagRepo struct {
	db *gorm.DB
}

func NewTagRepo() *TagRepo {
	return &TagRepo{
		db: database.GetDBClient(),
	}
}

func (repo *TagRepo) GetAllTags() []tagModel.Tag {
	var selectedTags []tagModel.Tag
	result := repo.db.Find(&selectedTags)

	if result.Error != nil {
		return nil
	}
	return selectedTags
}
