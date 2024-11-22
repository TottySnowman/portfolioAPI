package projectService

import (
	projectModel "portfolioAPI/project/models"
	project_repo "portfolioAPI/project/repos"
)

type ProjectService struct {
	repository *project_repo.Project_Repo
}

func NewProjectService() *ProjectService {
	return &ProjectService{
		project_repo.NewProjectRepo(),
	}
}

func (service *ProjectService) GetAllProjects() []projectModel.ProjectDisplay {
	return service.repository.GetAllProjects()
}

func (service *ProjectService) Insert(project projectModel.ProjectDisplay) error {
  databaseProject := GetDbProjectFromDisplay(project)
	return service.repository.Insert(&databaseProject)
}

func (service *ProjectService) Update(project projectModel.Project) error {
	return service.repository.Update(&project)
}

func (service *ProjectService) Delete(projectID int) error {
	return service.repository.Delete(projectID)
}

func GetDbProjectFromDisplay(display projectModel.ProjectDisplay) projectModel.Project {
	return projectModel.Project{
		Name:            display.Name,
		About:           display.About,
		Hidden:          display.Hidden,
		DevDate:         display.DevDate,
		DemoLink:        display.Demo_Link,
		GithubLink:      display.Github_Link,
		ProjectStatusID: uint(display.Status.StatusID),
	}
}
