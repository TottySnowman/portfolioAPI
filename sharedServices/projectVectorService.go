package sharedservices

import (
	projectModel "portfolioAPI/project/models"
	project_repo "portfolioAPI/project/repos"
)
type ProjectVectorService struct {
  projectRepo *project_repo.Project_Repo
}

func NewSharedService(projectRepo *project_repo.Project_Repo) *ProjectVectorService {
    return &ProjectVectorService{
    projectRepo: projectRepo,
  }
}

func (service *ProjectVectorService) GetAllProjects(includeHidden bool) []projectModel.ProjectDisplay {
	return service.projectRepo.GetAllProjects(includeHidden)
}

