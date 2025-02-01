package chatService

import "portfolioAPI/apiClients"

type ResponseService struct {
	client *apiClients.OpenAiClient
}

func NewResponseService(apiClient *apiClients.OpenAiClient) *ResponseService {
	return &ResponseService{
		client: apiClient,
	}
}

func (service *ResponseService) GetResponse(knowledgebase []string, prompt string) (string, error) {
	return service.client.GetSummaryResponse(knowledgebase, prompt)
}
