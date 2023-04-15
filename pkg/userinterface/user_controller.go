package userinterface

import (
	"net/http"
	"now-go-kon/pkg/application"
	"now-go-kon/pkg/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service application.UserService
}

func NewUserController(service application.UserService) *UserController {
	return &UserController{service: service}
}

type UserRequest struct {
	ID string `json:"id"`
}

func (c *UserController) GetProfileHandler(ctx *gin.Context) {
	var req UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	if err := c.service.User(ctx, domain.UserID(req.ID)); err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}
	ctx.Status(http.StatusOK)
}
