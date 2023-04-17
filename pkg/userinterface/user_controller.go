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
	ID int `json:"id"`
}

type userResponse struct {
	ID       int    `json:"id"`
	Username string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func newUserResponse(u *domain.User) userResponse {
	return userResponse{
		ID:       u.ID.Num(),
		Username: u.UserName.String(),
		Password: u.Password.String(),
		Email:    u.Email.String(),
	}
}

func (c *UserController) GetProfileHandler(ctx *gin.Context) {
	var req UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	user, err := c.service.User(ctx, domain.UserID(req.ID))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}
