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
