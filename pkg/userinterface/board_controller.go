package userinterface

import (
	"now-go-kon/pkg/application"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	service application.BoardService
}

func NewBoardController(service application.BoardService) *BoardController {
	return &BoardController{service: service}
}

type BoardRequest struct {
	ID int `json:"id"`
}

func (c *BoardController) CreateBoardHandler(ctx *gin.Context) {
	// var req BoardRequest
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
	// 	return
	// }

	// id := req.toParams()
	// user, err := c.service.User(ctx, id)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
	// 	return
	// }

	// res := UserProfileResponse(user)
	// ctx.JSON(http.StatusOK, res)
}
