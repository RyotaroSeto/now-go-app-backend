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
	board, err := c.service.Board(ctx, uParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	res := BoardCreateResponse(board)
	ctx.JSON(http.StatusOK, res)
}
