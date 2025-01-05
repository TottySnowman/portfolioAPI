package chatService

import (
	chatModel "portfolioAPI/chat/models"
	vectorRepo "portfolioAPI/chat/repos"
)

type VectorService struct {
	vectorRepo *vectorRepo.VectorRepo
}

func NewVectorService(vectorRepo *vectorRepo.VectorRepo) *VectorService {
  if !vectorRepo.DoesCollectionExist(){
    vectorRepo.CreateCollection()
  }

	return &VectorService{
		vectorRepo: vectorRepo,
	}
}

func (service *VectorService) UpsertVector(vector chatModel.FeatureExtractionResponse, promptModel chatModel.PromptModel) error {
	return service.vectorRepo.UpsertVector(vector, promptModel)
}

func (service *VectorService) ResetDatabase(syncModel *chatModel.SyncModel) error {
	if !syncModel.ResetProject && !syncModel.ResetPersonal {
		return nil
	}

	if syncModel.ResetProject && syncModel.ResetPersonal {
		return service.vectorRepo.FullResetDatabase()
	}

	if syncModel.ResetProject {
		return service.vectorRepo.ResetProject()
	}

	if syncModel.ResetPersonal {
		return service.vectorRepo.ResetPersonal()
	}

	return nil
}
