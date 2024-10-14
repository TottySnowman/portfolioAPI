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

func(service *ProjectService) Insert(project projectModel.Project) error{
  return service.repository.Insert(&project)
}

func(service *ProjectService) Update(project projectModel.Project) error{
  return service.repository.Update(&project)
}
