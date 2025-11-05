package tagService

import (
	tagModel "portfolioAPI/tag/models"
	tagRepo "portfolioAPI/tag/repos"
)

type TagService struct {
	repo *tagRepo.TagRepo
}

func NewTagService(tagRepo *tagRepo.TagRepo) *TagService {
	return &TagService{
		repo: tagRepo,
	}
}

func (service *TagService) GetAllTags() []tagModel.JsonTag {
	return service.repo.GetAllTags()
}

func (service *TagService) GetTechStack() tagModel.TechstackResponse {
	var backendItems []tagModel.TechstackItemResponse
	var frontendItems []tagModel.TechstackItemResponse

	techStackItems := service.repo.GetTechStack()
	for _, item := range techStackItems {
		if item.IsBackend {
			backendItems = append(backendItems, service.mapTechstackItemToTechStackItemResponse(item))
			continue
		}
		frontendItems = append(frontendItems, service.mapTechstackItemToTechStackItemResponse(item))
	}

	return tagModel.TechstackResponse{
		Backend:  backendItems,
		Frontend: frontendItems,
	}
}

func (service *TagService) mapTechstackItemToTechStackItemResponse(techStackItem tagModel.TechstackItem) tagModel.TechstackItemResponse {
	return tagModel.TechstackItemResponse{
		DisplayName: techStackItem.Tag,
		IconName:    techStackItem.Icon,
	}
}

func (service *TagService) Insert(tag tagModel.Tag) (tagModel.JsonTag, error) {
	createdTag, err := service.repo.Insert(&tag)
	convertedTag := service.ConvertTagToDisplayTag(createdTag)
	return convertedTag, err
}

func (service *TagService) Update(tag tagModel.JsonTag) (tagModel.JsonTag, error) {
	convertedTag := service.ConvertDisplayTagToDbTag(tag)
	updatedTag, err := service.repo.Update(&convertedTag)
	convertedUpdatedTag := service.ConvertTagToDisplayTag(updatedTag)
	return convertedUpdatedTag, err
}

func (service *TagService) Delete(tagID int) error {
	return service.repo.Delete(tagID)
}

func (service *TagService) ConvertTagToDisplayTag(dbTag *tagModel.Tag) tagModel.JsonTag {
	return tagModel.JsonTag{
		TagId: int(dbTag.ID),
		Icon:  dbTag.Icon,
		Tag:   dbTag.Tag,
	}
}

func (service *TagService) ConvertDisplayTagToDbTag(displayTag tagModel.JsonTag) tagModel.Tag {
	return tagModel.Tag{
		ID:   uint(displayTag.TagId),
		Icon: displayTag.Icon,
		Tag:  displayTag.Tag,
	}
}
