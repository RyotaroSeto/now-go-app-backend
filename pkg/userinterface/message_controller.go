package userinterface

import (
	"now-go-kon/pkg/application"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	service application.MessageService
}

func NewMessageController(service application.MessageService) *MessageController {
	return &MessageController{service: service}
}

func (c *MessageController) GetMessageHandler(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, res)
}

func (c *MessageController) GetMessageScrollHandler(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, res)
}

func (c *MessageController) CreateMessageHandler(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, res)
}
