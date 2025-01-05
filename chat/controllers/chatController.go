package chatController

import (
	"net/http"
	chatModel "portfolioAPI/chat/models"
	chatService "portfolioAPI/chat/services"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	embeddingService *chatService.EmbeddingService
	vectorService    *chatService.VectorService
}

func NewChatController(embeddingService *chatService.EmbeddingService, vectorService *chatService.VectorService) *ChatController {
	return &ChatController{
		embeddingService: embeddingService,
		vectorService:    vectorService,
	}
}

func (con *ChatController) Test(context *gin.Context) {

	var prompt *chatModel.PromptModel
	if err := context.ShouldBindJSON(&prompt); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	print(prompt)
  err, vector := con.embeddingService.GetVectorByText(prompt.Prompt)
  if err != nil{
    panic(err.Error())
  }

  //con.vectorService.CreateCollectionIfNeeded()
  con.vectorService.InsertVector(vector, prompt.Prompt)

}
