package projectService

import (
	"os"
	fileServices "portfolioAPI/fileUpload/services"
	projectModel "portfolioAPI/project/models"
	project_repo "portfolioAPI/project/repos"
	tagModel "portfolioAPI/tag/models"
	tagService "portfolioAPI/tag/services"
	"strings"
)

type ProjectUpdateListener interface {
	OnProjectUpdated(project projectModel.ProjectDisplay)
	OnProjectDeleted(projectId int)
}

type ProjectService struct {
	repository      *project_repo.Project_Repo
	tagService      *tagService.TagService
	fileService     *fileServices.FileService
	updateListeners []ProjectUpdateListener
}

func NewProjectService(projectRepo *project_repo.Project_Repo,
	tagService *tagService.TagService,
	fileService *fileServices.FileService) *ProjectService {
	return &ProjectService{
		repository:      projectRepo,
		tagService:      tagService,
		fileService:     fileService,
		updateListeners: []ProjectUpdateListener{},
	}
}
func (service *ProjectService) RegisterListener(listener ProjectUpdateListener) {
	service.updateListeners = append(service.updateListeners, listener)
}

func (service *ProjectService) notifyProjectUpdated(project projectModel.ProjectDisplay) {
	for _, listener := range service.updateListeners {
		listener.OnProjectUpdated(project)
	}
}

func (service *ProjectService) notifyProjectDeleted(projectId int) {
	for _, listener := range service.updateListeners {
		listener.OnProjectDeleted(projectId)
	}
}

func (service *ProjectService) GetAllProjects(includeHidden bool, languageCode string) []projectModel.ProjectDisplay {
	return service.repository.GetAllProjects(includeHidden, languageCode)
}

func (service *ProjectService) Insert(project projectModel.ProjectDisplay, languageCode string) (*projectModel.ProjectDisplay, error) {
	databaseProject := GetDbProjectFromDisplay(project, languageCode)
	_, insertError := service.repository.Insert(&databaseProject)

	if insertError != nil {
		return nil, insertError
	}

	service.insertIntoProjectTag(project.Tags, databaseProject.ID)

	mappedProject, err := service.GetProjectById(databaseProject.ID, true)
	if err != nil {
		return nil, err
	}

	if !mappedProject.Hidden {
		service.notifyProjectUpdated(*mappedProject)
	}

	return mappedProject, nil
}

func (service *ProjectService) Update(project projectModel.ProjectDisplay, languageCode string) (*projectModel.ProjectDisplay, error) {
	databaseProject := GetDbProjectFromDisplay(project, languageCode)
	_, error := service.repository.Update(&databaseProject)

	if error != nil {
		return nil, error
	}

	service.insertIntoProjectTag(project.Tags, project.ProjectID)

	mappedProject, err := service.GetProjectById(databaseProject.ID, true)
	if err != nil {
		return nil, err
	}

	if !mappedProject.Hidden {
		service.notifyProjectUpdated(*mappedProject)
	}
	return mappedProject, nil
}

func (service *ProjectService) insertIntoProjectTag(projectTags []tagModel.JsonTag, projectId int) {
	convertedTags := service.convertDisplayTagsToDbTags(projectTags)
	tagIds := getTagIds(convertedTags)

	error := service.repository.InsertIntoProjectTags(projectId, tagIds)

	if error != nil {
		// TODO handling
		println("Error in insert")
	}

}

func getTagIds(tags []tagModel.Tag) []int {
	var tagIds []int
	for _, tag := range tags {
		tagIds = append(tagIds, int(tag.ID))
	}

	return tagIds
}

func (service *ProjectService) Delete(projectID int) error {
	existingProject, err := service.GetProjectById(projectID, false)
	if err != nil {
		return err
	}

	err = service.repository.Delete(projectID)
	if err != nil {
		return err
	}

	if !existingProject.Hidden {
		service.notifyProjectDeleted(existingProject.ProjectID)
	}

	existingProject.Logo_Path = service.removeUrlPrefix(existingProject.Logo_Path)
	return service.fileService.HandleFileDelete("/logo", existingProject.Logo_Path)
}

func (service *ProjectService) removeUrlPrefix(url string) string {
	return strings.TrimPrefix(url, os.Getenv("API_ENDPOINT_URL"))
}

func (service *ProjectService) GetProjectById(projectId int, includeHidden bool) (*projectModel.ProjectDisplay, error) {
	return service.repository.GetProjectById(projectId, includeHidden)
}

func (service *ProjectService) convertDisplayTagsToDbTags(projectTags []tagModel.JsonTag) []tagModel.Tag {
	var convertedTags []tagModel.Tag
	for _, tag := range projectTags {
		convertedTags = append(convertedTags, service.tagService.ConvertDisplayTagToDbTag(tag))
	}

	return convertedTags
}

func GetDbProjectFromDisplay(display projectModel.ProjectDisplay, languageCode string) projectModel.Project {
	return projectModel.Project{
		ID:              display.ProjectID,
		Name:            display.Name,
		About:           display.About,
		Hidden:          display.Hidden,
		DevDate:         display.DevDate,
		DemoLink:        display.Demo_Link,
		GithubLink:      display.Github_Link,
		ProjectStatusID: uint(display.Status.StatusID),
		LogoPath:        display.Logo_Path,
		LanguageCode:    display.LanguageCode,
	}
}
