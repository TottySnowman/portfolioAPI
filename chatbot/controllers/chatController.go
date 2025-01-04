package chatController

import (
	"github.com/gin-gonic/gin"
)

type ChatController struct{

}

func NewChatController() *ChatController{
  return &ChatController{}
}

func(con *ChatController) Test(context *gin.Context){

}
