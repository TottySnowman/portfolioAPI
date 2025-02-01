package chatService

import (
	"github.com/coder/websocket"
	"github.com/gin-gonic/gin"
)

type WsService struct {}

func NewWsService() *WsService {
	return &WsService{}
}

func(service *WsService) GetWebsocketConnection(ctx *gin.Context)*websocket.Conn{

	acceptOptions := websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	}
	conn, err := websocket.Accept(ctx.Writer, ctx.Request, &acceptOptions)
	if err != nil {
		return nil
	}

  return conn
}
