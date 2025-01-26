package knowledgeController

import (
	"net/http"
	chatModel "portfolioAPI/chat/models"
	chatService "portfolioAPI/chat/services"

	"github.com/gin-gonic/gin"
)

type KnowledgeController struct {
	vectorService    *chatService.VectorService
	embeddingService *chatService.EmbeddingService
}

func NewKnowledgeController(vectorService *chatService.VectorService) *KnowledgeController {
	return &KnowledgeController{
		vectorService: vectorService,
	}
}

func (con *KnowledgeController) InsertText(ctx *gin.Context) {
	var prompt *chatModel.PromptModel
	if err := ctx.ShouldBindJSON(&prompt); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	vector, err := con.embeddingService.GetVectorByText(prompt.Prompt)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	err = con.vectorService.UpsertText(vector, prompt.Prompt, prompt.PointId) // TODO get the inserted point back and return it

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	}
}
