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

	_, err := con.embeddingService.GetVectorByText(prompt.Prompt)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	}
}

func (con *ChatController) FullSync(context *gin.Context) {
	syncSettings := &chatModel.SyncModel{
		ResetProject:  true,
		ResetPersonal: true,
	}

	if err := con.vectorService.ResetDatabase(syncSettings); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	con.vectorService.InsertProjectsAsync()
}

func (con *ChatController) Chat(context *gin.Context) {
	var prompt *chatModel.PromptModel

	if err := context.ShouldBindJSON(&prompt); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	message, err := con.vectorService.GetChatMessage(prompt)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

  context.IndentedJSON(http.StatusOK, gin.H{"message": message})
}

func (con *ChatController) Sync(context *gin.Context) {
	var syncSettings *chatModel.SyncModel
	if err := context.ShouldBindJSON(&syncSettings); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	if err := con.vectorService.ResetDatabase(syncSettings); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

}
