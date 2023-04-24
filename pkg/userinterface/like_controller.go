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

type GetLikeRequest struct {
	UserID int `json:"user_id"`
}

// func LikedGetResponse(us []*domain.Like) []GetBoardResponse {
// 	var br []GetBoardResponse
// 	for _, v := range us {
// 		br = append(br, GetBoardResponse{
// 			BoardID:     v.ID.Num(),
// 			UserID:      v.UserID.Num(),
// 			Body:        v.Body.String(),
// 			CreatedDate: v.CreatedDate,
// 		})
// 	}
// 	return br
// }

func (c *LikeController) GetLikeHandler(ctx *gin.Context) {
	var req LikeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	uID := domain.UserID(req.UserID)
	_, err := c.service.LikeGet(ctx, uID)
	// likes, err := c.service.LikeGet(ctx, uID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	// res := LikedGetResponse(likes)
	// ctx.JSON(http.StatusOK, res)
	ctx.Status(http.StatusOK)
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
