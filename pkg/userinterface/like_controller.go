package userinterface

import (
	"net/http"
	"now-go-kon/pkg/application"
	"now-go-kon/pkg/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	service application.LikeService
}

func NewLikeController(service application.LikeService) *LikeController {
	return &LikeController{service: service}
}

type GetLikeRequest struct {
	UserID int `json:"user_id"`
}

type GetLikedResponse struct {
	LikedUserID int       `json:"liked_user_id"`
	LikedDate   time.Time `json:"liked_date"`
	MessageBody string    `json:"message_body"`
}

func LikedGetResponse(us []*domain.Like) []GetLikedResponse {
	var lr []GetLikedResponse
	for _, v := range us {
		lr = append(lr, GetLikedResponse{
			LikedUserID: v.LikedUserID.Num(),
			LikedDate:   v.LikedDate,
			MessageBody: v.MessageBody.String(),
		})
	}
	return lr
}

func (c *LikeController) GetLikeHandler(ctx *gin.Context) {
	var req LikeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	uID := domain.UserID(req.UserID)
	likes, err := c.service.LikeGet(ctx, uID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	res := LikedGetResponse(likes)
	ctx.JSON(http.StatusOK, res)
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
