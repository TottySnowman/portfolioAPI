package chatController

import (
	"context"
	"fmt"
	"net/http"
	chatModel "portfolioAPI/chat/models"
	chatService "portfolioAPI/chat/services"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	embeddingService *chatService.EmbeddingService
	vectorService    *chatService.VectorService
	wsService        *chatService.WsService
}

func NewChatController(embeddingService *chatService.EmbeddingService,
	vectorService *chatService.VectorService,
	wsService *chatService.WsService) *ChatController {
	return &ChatController{
		embeddingService: embeddingService,
		vectorService:    vectorService,
		wsService:        wsService,
	}
}

func (con *ChatController) InsertKnowledge(context *gin.Context) {
	var prompt *chatModel.PromptModel
	if err := context.ShouldBindJSON(&prompt); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "invalid"})
		return
	}

	vector, err := con.embeddingService.GetVectorByText(prompt.Prompt)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	err = con.vectorService.UpsertText(vector, prompt.Prompt, prompt.PointId) // TODO get the inserted point back and return it

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

func (con *ChatController) CreateWsConnection(cxt *gin.Context) {
	wsConnection := con.wsService.GetWebsocketConnection(cxt)
	if wsConnection == nil {
		cxt.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for {
		typ, data, err := wsConnection.Read(context.Background())
		if err != nil {
			println(err)
			break
		}

		if !con.embeddingService.IsModelReady() {
			wsConnection.Write(context.Background(), typ, []byte("Model is starting..."))
			go func() {
				con.embeddingService.StartModel()
				wsConnection.Write(context.Background(), typ, []byte("Model started successfully!"))
			}()
			continue
		} else {
			wsConnection.Write(context.Background(), typ, []byte("Model is already running!"))
		}

		if string(data) == "ping" {
			err = wsConnection.Write(context.Background(), typ, []byte("pong"))
		} else {
			err = wsConnection.Write(context.Background(), typ, []byte(fmt.Sprintf("Hello, %s", string(data))))
		}

		if err != nil {
			println(err)
			break
		}
	}

}
