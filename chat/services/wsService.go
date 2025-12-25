package chatService

import (
	"os"

	"github.com/coder/websocket"
	"github.com/gin-gonic/gin"
)

type WsService struct{}

func NewWsService() *WsService {
	return &WsService{}
}

func (service *WsService) GetWebsocketConnection(ctx *gin.Context) *websocket.Conn {
  ws_cors0 := os.Getenv("WS_CORS_ORIGIN0")
  ws_cors1 := os.Getenv("WS_CORS_ORIGIN1")
	acceptOptions := websocket.AcceptOptions{
    OriginPatterns: []string{ws_cors0, ws_cors1},
	}

	conn, err := websocket.Accept(ctx.Writer, ctx.Request, &acceptOptions)
	if err != nil {
		return nil
	}

	return conn
}
