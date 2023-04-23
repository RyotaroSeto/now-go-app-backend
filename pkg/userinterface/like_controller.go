package userinterface

import (
	"net/http"
	"now-go-kon/pkg/application"
	"now-go-kon/pkg/domain"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	service application.LikeService
}

func NewLikeController(service application.LikeService) *LikeController {
	return &LikeController{service: service}
}

type LikeRequest struct {
	UserID      int    `json:"user_id"`
	LikedUserID int    `json:"liked_user_id"`
	MessageBody string `json:"message_body"`
}

func (r *LikeRequest) toParams() *domain.Like {
	return &domain.Like{
		UserID:      domain.UserID(r.UserID),
		LikedUserID: domain.UserID(r.LikedUserID),
		MessageBody: domain.MessageBody(r.MessageBody),
	}
}

func (c *LikeController) CreateLikeHandler(ctx *gin.Context) {
	var req LikeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	uParam := req.toParams()
	if err := c.service.LikeCreate(ctx, uParam); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	ctx.Status(http.StatusOK)
}
