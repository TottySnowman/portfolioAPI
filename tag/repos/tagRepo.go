package tagRepo

import (
	"errors"
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

func (repo *TagRepo) GetAllTags() []tagModel.JsonTag {
	var selectedTags []tagModel.Tag
	result := repo.db.Find(&selectedTags)

	if result.Error != nil {
		return nil
	}

  var jsonTags []tagModel.JsonTag
  for _, tag := range selectedTags {
    jsonTags = append(jsonTags, tagModel.JsonTag{
			TagId:   int(tag.ID),
			Tag:     tag.Tag,
			Icon: tag.Icon,
		})
  }
  
	return jsonTags
}

func (repo *TagRepo) Insert(tagToCreate *tagModel.Tag) error {
	result := repo.db.Create(&tagToCreate)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *TagRepo) Update(tagToUpdate *tagModel.Tag) error {
	var dbTag = tagModel.Tag{Model: gorm.Model{ID: tagToUpdate.ID}}
	existingTag := repo.db.First(&dbTag)

	if existingTag.Error != nil {
		return errors.New("Tag not found")
	}

	updateTag := repo.db.Model(&dbTag).Select("*").Omit("CreatedAt").Updates(tagToUpdate)

	if updateTag.Error != nil {
		return errors.New(updateTag.Error.Error())
	}

	return nil
}

func (repo *TagRepo) Delete(tagID int) error {

	var dbTag = tagModel.Tag{Model: gorm.Model{ID: uint(tagID)}}
	existingTag := repo.db.First(&dbTag)

	if existingTag.Error != nil {
		return errors.New("Tag not found")
	}

	result := repo.db.Delete(&dbTag)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
