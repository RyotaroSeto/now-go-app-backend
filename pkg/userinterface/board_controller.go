package userinterface

import (
	"net/http"
	"now-go-kon/pkg/application"
	"now-go-kon/pkg/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	service application.BoardService
}

func NewBoardController(service application.BoardService) *BoardController {
	return &BoardController{service: service}
}

type BoardGetRequest struct {
	Gender string `form:"gender"`
}

type GetBoardResponse struct {
	BoardID     int       `json:"board_id"`
	UserID      int       `json:"user_id"`
	Body        string    `json:"body"`
	CreatedDate time.Time `json:"created_date"`
}

func BoardGetResponse(us []*domain.Board) []GetBoardResponse {
	var br []GetBoardResponse
	for _, v := range us {
		br = append(br, GetBoardResponse{
			BoardID:     v.ID.Num(),
			UserID:      v.UserID.Num(),
			Body:        v.Body.String(),
			CreatedDate: v.CreatedDate,
		})
	}
	return br
}

func (c *BoardController) GetBoardHandler(ctx *gin.Context) {
	var req BoardGetRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	dGender := domain.Gender(req.Gender)
	boards, err := c.service.BoardGet(ctx, dGender)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	res := BoardGetResponse(boards)
	ctx.JSON(http.StatusOK, res)
}

type ScrollRequest struct {
	Gender  string `form:"gender"`
	BoardID int    `form:"board_id"`
}

func (c *BoardController) GetScrollHandler(ctx *gin.Context) {
	var req ScrollRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	dGender := domain.Gender(req.Gender)
	lastBordID := domain.BoardID(req.BoardID)
	nextBoards, err := c.service.ScrollBoardGet(ctx, dGender, lastBordID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	res := BoardGetResponse(nextBoards)
	ctx.JSON(http.StatusOK, res)
}

type BoardRequest struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

type BoardResponse struct {
	ID int `json:"id"`
}

func BoardCreateResponse(u *domain.Board) BoardResponse {
	return BoardResponse{
		ID: u.UserID.Num(),
	}
}

func (r *BoardRequest) toParams() *domain.Board {
	return &domain.Board{
		UserID: domain.UserID(r.ID),
		Body:   domain.Body(r.Body),
	}
}

func (c *BoardController) CreateBoardHandler(ctx *gin.Context) {
	var req BoardRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	uParam := req.toParams()
	board, err := c.service.BoardCreate(ctx, uParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	res := BoardCreateResponse(board)
	ctx.JSON(http.StatusOK, res)
}

type BoardDeleteRequest struct {
	ID int `json:"id"`
}

func (c *BoardController) DeleteBoardHandler(ctx *gin.Context) {
	var req BoardDeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	err := c.service.BoardDelete(ctx, domain.BoardID(req.ID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	ctx.Status(http.StatusOK)
}
