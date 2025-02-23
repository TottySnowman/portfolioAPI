package statusService

import (
	statusModels "portfolioAPI/status/models"
	statusRepo "portfolioAPI/status/repos"
)


type StatusService struct {
	repo *statusRepo.StatusRepo
}

func NewStatusService(statusRepo *statusRepo.StatusRepo) *StatusService {
	return &StatusService{
		repo: statusRepo,
	}
}

func(service *StatusService) GetAllStatus() []statusModels.ProjectStatusDisplay{
  return service.repo.GetAllStatus()
}
