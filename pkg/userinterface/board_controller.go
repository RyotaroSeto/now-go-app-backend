package userinterface

import (
	"net/http"
	"now-go-kon/pkg/application"
	"now-go-kon/pkg/domain"

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
	ID   int    `json:"id"`
	Body string `json:"body"`
}

func BoardGetResponse(us []*domain.Board) []GetBoardResponse {
	var br []GetBoardResponse
	for _, v := range us {
		br = append(br, GetBoardResponse{ID: v.UserID.Num(), Body: v.Body.String()})
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
	board, err := c.service.BoardGet(ctx, dGender)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	res := BoardGetResponse(board)
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

	bID := domain.BoardID(req.ID)
	_, err := c.service.BoardDelete(ctx, bID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	ctx.Status(http.StatusOK)
}
