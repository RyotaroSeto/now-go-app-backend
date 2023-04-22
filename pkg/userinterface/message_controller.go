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

func (c *MessageController) CreateMessageHandler(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, res)
}
