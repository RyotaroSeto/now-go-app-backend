package userinterface

import (
	"log"
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

func (r *UserRequest) toParams() domain.UserID {
	return domain.UserID((r.ID))
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func UserProfileResponse(u *domain.User) UserResponse {
	return UserResponse{
		ID:       u.ID.Num(),
		Username: u.UserName.String(),
		Password: u.Password.String(),
		Email:    u.Email.String(),
	}
}

func (c *UserController) GetProfileHandler(ctx *gin.Context) {
	var req UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	id := req.toParams()
	user, err := c.service.User(ctx, id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	res := UserProfileResponse(user)
	ctx.JSON(http.StatusOK, res)
}
