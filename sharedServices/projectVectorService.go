package sharedservices

import (
	vectorRepo "portfolioAPI/chat/repos"
	project_repo "portfolioAPI/project/repos"
)
type ProjectVectorService struct {
  vectorRepo *vectorRepo.VectorRepo
  project_repo *project_repo.Project_Repo
}

func NewSharedService() *ProjectVectorService {
    return &ProjectVectorService{}
}


