package chatService

import (
	"portfolioAPI/apiClients"
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

func (service *EmbeddingService) StartModel() {
	if isModelStarting {
		return
	}
	isModelStarting = true

	for {
		_, err := service.huggingFaceApiClient.GetVectorByText("Are you alive?")

		if err == nil {
			modelReadyState = true
      isModelStarting = false
      go service.MonitorModelHealth()
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (service *EmbeddingService) IsModelReady() bool {
	if modelReadyState == true {
		return true
	}

	return false
}

func (service *EmbeddingService) IsModelStarting() bool {
	return isModelStarting
}

func (service *EmbeddingService) MonitorModelHealth() {
    for {
        _, err := service.huggingFaceApiClient.GetVectorByText("Are you alive?")
        if err == nil {
            modelReadyState = true
        } else {
            modelReadyState = false
        }
        time.Sleep(3600 * time.Second)
    }
}
