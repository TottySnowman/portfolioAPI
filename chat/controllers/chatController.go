package chatController

import (
	"context"
	"fmt"
	"net/http"
	chatModel "portfolioAPI/chat/models"
	chatService "portfolioAPI/chat/services"

	"github.com/coder/websocket"
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

var modelReady = false
func (con *ChatController) CreateWsConnection(cxt *gin.Context) {
	acceptOptions := websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	}
	conn, err := websocket.Accept(cxt.Writer, cxt.Request, &acceptOptions)
	if err != nil {
		http.Error(cxt.Writer, "Failed to upgrade WebSocket connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "closing connection")

	for {
		typ, data, err := conn.Read(context.Background())
		if err != nil {
			break
		}

		if string(data) == "ping" {
			err = conn.Write(context.Background(), typ, []byte("pong"))
		} else {
			err = conn.Write(context.Background(), typ, []byte(fmt.Sprintf("Hello, %s", string(data))))
		}

		if err != nil {
			break
		}
	}
}
