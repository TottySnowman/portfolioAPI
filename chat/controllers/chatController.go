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

func (con *ChatController) Upsert(context *gin.Context) {
	var prompt *chatModel.PromptModel
	if err := context.ShouldBindJSON(&prompt); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	err, vector := con.embeddingService.GetVectorByText(prompt.Prompt)
	if err != nil {
		panic(err.Error())
	}

	err = con.vectorService.UpsertVector(vector, *prompt)

	context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	return
}

func (con *ChatController) FullSync(context *gin.Context) {
	var syncSettings *chatModel.SyncModel
	if err := context.ShouldBindJSON(&syncSettings); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}
}
func (con *ChatController) Sync(context *gin.Context) {
	var syncSettings *chatModel.SyncModel
	if err := context.ShouldBindJSON(&syncSettings); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}


}
