package tagService

import (
	tagModel "portfolioAPI/tag/models"
	tagRepo "portfolioAPI/tag/repos"
)

type TagService struct{
  repo *tagRepo.TagRepo
}

func NewTagService() *TagService{
  return &TagService{
    repo: tagRepo.NewTagRepo(),
  }
}

func(service *TagService) GetAllTags() []tagModel.Tag{
  return service.repo.GetAllTags()
}
func (service *TagService) Insert(tag tagModel.Tag) error {
	return service.repo.Insert(&tag)
}

func (service *TagService) Update(tag tagModel.Tag) error {
	return service.repo.Update(&tag)
}

func (service *TagService) Delete(tagID int) error {
	return service.repo.Delete(tagID)
}

