package chatService

import (
	"portfolioAPI/chat/apiClients"
	chatModel "portfolioAPI/chat/models"
)

type EmbeddingService struct {
	huggingFaceApiClient *apiClients.HuggingFaceClient
}

func NewEmbeddingService(huggingFaceApiClient *apiClients.HuggingFaceClient) *EmbeddingService {
	return &EmbeddingService{
		huggingFaceApiClient: huggingFaceApiClient,
	}
}

func (service *EmbeddingService) GetVectorByText(text string) (chatModel.FeatureExtractionResponse, error) {
	return service.huggingFaceApiClient.GetVectorByText(text)
}
