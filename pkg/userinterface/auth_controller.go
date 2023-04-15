package userinterface

import (
	"net/http"
	"now-go-kon/pkg/application"
	"now-go-kon/pkg/domain"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service application.AuthService
}

func NewAuthController(service application.AuthService) *AuthController {
	return &AuthController{service: service}
}

type AuthRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func (c *AuthController) PasswordAuthHandler(ctx *gin.Context) {
	var req AuthRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	if err := c.service.Auth(ctx, domain.UserID(req.ID), domain.Password(req.Password)); err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}
	ctx.Status(http.StatusOK)
}
