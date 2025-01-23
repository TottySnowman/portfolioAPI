package chatService

import (
	"portfolioAPI/chat/apiClients"
	chatModel "portfolioAPI/chat/models"
	"time"
)

type EmbeddingService struct {
	huggingFaceApiClient *apiClients.HuggingFaceClient
}

var modelReadyState = false
var isModelStarting = false
func NewEmbeddingService(huggingFaceApiClient *apiClients.HuggingFaceClient) *EmbeddingService {
	return &EmbeddingService{
		huggingFaceApiClient: huggingFaceApiClient,
	}
}

func (service *EmbeddingService) GetVectorByText(text string) (chatModel.FeatureExtractionResponse, error) {
	return service.huggingFaceApiClient.GetVectorByText(text)
}

func (service *EmbeddingService) StartModel(){
  time.Sleep(10 * time.Second)
  modelReadyState = true
}

func (service *EmbeddingService) IsModelReady() bool{
  if modelReadyState == true{
    return true
  }

  return false
}

func (service *EmbeddingService) IsModelStarting() bool{
  return isModelStarting
}
