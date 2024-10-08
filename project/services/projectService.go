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
