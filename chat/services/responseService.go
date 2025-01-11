package chatService
type ResponseService struct {
}

func NewResponseService() *ResponseService {
	return &ResponseService{}
}

func(service *ResponseService) GetResponse(knowledgebase []string) (*string, error){
  return nil, nil
}
