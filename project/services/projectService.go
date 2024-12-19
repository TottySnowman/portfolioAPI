package projectService

import (
	projectModel "portfolioAPI/project/models"
	project_repo "portfolioAPI/project/repos"
	tagModel "portfolioAPI/tag/models"
	tagService "portfolioAPI/tag/services"
)

type ProjectService struct {
	repository *project_repo.Project_Repo
	tagService *tagService.TagService
}

func NewProjectService(projectRepo *project_repo.Project_Repo, tagService *tagService.TagService) *ProjectService {
	return &ProjectService{
		repository: projectRepo,
		tagService: tagService,
	}
}

func (service *ProjectService) GetAllProjects() []projectModel.ProjectDisplay {
	return service.repository.GetAllProjects()
}

func (service *ProjectService) Insert(project projectModel.ProjectDisplay) error {
	databaseProject := GetDbProjectFromDisplay(project)
  insertError := service.repository.Insert(&databaseProject)
  if insertError != nil{
    return insertError
  }

  println("BING BONG!")
  println("NEW ID OF PROJECT: ", databaseProject.ID)
  service.insertIntoProjectTag(project.Tags, databaseProject.ID)
  return nil
}

func (service *ProjectService) Update(project projectModel.ProjectDisplay) error {
	databaseProject := GetDbProjectFromDisplay(project)
  error := service.repository.Update(&databaseProject)
  if error != nil{
    return error
  }

  service.insertIntoProjectTag(project.Tags, project.ProjectID)
  return nil
}

func (service *ProjectService) Delete(projectID int) error {
	return service.repository.Delete(projectID)
}

func (service *ProjectService) insertIntoProjectTag(projectTags []tagModel.JsonTag, projectId int) {
	convertedTags := service.convertDisplayTagsToDbTags(projectTags)
  tagIds := getTagIds(convertedTags)

  error := service.repository.InsertIntoProjectTags(projectId, tagIds)

  if error != nil{
    // TODO handling
    println("Error in insert")
  }

}

func (service *ProjectService)convertDisplayTagsToDbTags(projectTags []tagModel.JsonTag) []tagModel.Tag {
	var convertedTags []tagModel.Tag
	for _, tag := range projectTags {
    println(tag.TagId)
		convertedTags = append(convertedTags, service.tagService.ConvertDisplayTagToDbTag(tag))
	}

	return convertedTags
}

func getTagIds(tags []tagModel.Tag) []int{
  var tagIds []int
  for _, tag := range tags{
    tagIds = append(tagIds, int(tag.ID))
  }

  return tagIds
}

func GetDbProjectFromDisplay(display projectModel.ProjectDisplay) projectModel.Project {
	return projectModel.Project{
    ID:             display.ProjectID,
		Name:            display.Name,
		About:           display.About,
		Hidden:          display.Hidden,
		DevDate:         display.DevDate,
		DemoLink:        display.Demo_Link,
		GithubLink:      display.Github_Link,
		ProjectStatusID: uint(display.Status.StatusID),
	}
}
