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
	ID int `form:"id"`
}

type UserResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"user_name"`
	Status      int    `json:"status"`
	Email       string `json:"email"`
	DateOfBirth int    `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Residence   string `json:"residence"`
	Occupation  string `json:"occupation"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
}

func UserProfileResponse(u *domain.User) UserResponse {
	return UserResponse{
		ID:          u.ID.Num(),
		Username:    u.UserName.String(),
		Status:      u.Status.Num(),
		Email:       u.Email.String(),
		DateOfBirth: u.UsersDetails.DateOfBirth.Num(),
		Gender:      u.UsersDetails.Gender.String(),
		Residence:   u.UsersDetails.Residence.String(),
		Occupation:  u.UsersDetails.Occupation.String(),
		Height:      u.UsersDetails.Height.Num(),
		Weight:      u.UsersDetails.Weight.Num(),
	}
}

func (c *UserController) GetProfileHandler(ctx *gin.Context) {
	var req UserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	id := domain.UserID(req.ID)
	user, err := c.service.User(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	res := UserProfileResponse(user)
	ctx.JSON(http.StatusOK, res)
}

type UserUpdateRequest struct {
	ID          int    `json:"id"`
	DateOfBirth int    `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Residence   string `json:"residence"`
	Occupation  string `json:"occupation"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
}

type UserDetailResponse struct {
	UserID      int    `json:"user_id"`
	DateOfBirth int    `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Residence   string `json:"residence"`
	Occupation  string `json:"occupation"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
}

func UserUpdateResponse(u *domain.UsersDetails) UserDetailResponse {
	return UserDetailResponse{
		UserID:      u.UserID.Num(),
		DateOfBirth: u.DateOfBirth.Num(),
		Gender:      u.Gender.String(),
		Residence:   u.Residence.String(),
		Occupation:  u.Occupation.String(),
		Height:      u.Height.Num(),
		Weight:      u.Weight.Num(),
	}
}

func (r *UserUpdateRequest) toParams() *domain.UsersDetails {
	return &domain.UsersDetails{
		UserID:      domain.UserID(r.ID),
		DateOfBirth: domain.DateOfBirth(r.DateOfBirth),
		Gender:      domain.Gender(r.Gender),
		Residence:   domain.Residence(r.Residence),
		Occupation:  domain.Occupation(r.Occupation),
		Height:      domain.Height(r.Height),
		Weight:      domain.Weight(r.Weight),
	}
}

func (c *UserController) UpdateProfileHandler(ctx *gin.Context) {
	var req UserUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	uParam := req.toParams()
	updateUser, err := c.service.UserUpdate(ctx, uParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	res := UserUpdateResponse(updateUser)
	ctx.JSON(http.StatusOK, res)
}
