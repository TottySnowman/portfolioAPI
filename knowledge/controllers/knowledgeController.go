package knowledgeController

import (
	"net/http"
	chatModel "portfolioAPI/chat/models"
	chatService "portfolioAPI/chat/services"
	knowledgeModels "portfolioAPI/knowledge/models"

	"github.com/gin-gonic/gin"
)

type KnowledgeController struct {
	vectorService    *chatService.VectorService
	embeddingService *chatService.EmbeddingService
}

func NewKnowledgeController(vectorService *chatService.VectorService, embeddingService *chatService.EmbeddingService) *KnowledgeController {
	return &KnowledgeController{
		vectorService:    vectorService,
		embeddingService: embeddingService,
	}
}

func (con *KnowledgeController) UpsertText(ctx *gin.Context) {
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

	createdPoint, err := con.vectorService.UpsertText(vector, prompt.Prompt, prompt.PointId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	}
	ctx.IndentedJSON(http.StatusCreated, createdPoint)
}

func (con *KnowledgeController) DeleteSinglePoint(ctx *gin.Context) {
	var deleteModel *knowledgeModels.DeleteModel
	if err := ctx.ShouldBindJSON(&deleteModel); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	if err := con.vectorService.DeleteSinglePoint(deleteModel.PointId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	ctx.Status(http.StatusOK)
}

func (con *KnowledgeController) GetKnowledgeBase(ctx *gin.Context) {
	knowledgeBase := con.vectorService.GetFullKnowledgeBase()
	ctx.IndentedJSON(http.StatusOK, knowledgeBase)
}
