package userinterface

import (
	"net/http"
	"now-go-kon/pkg/application"
	"now-go-kon/pkg/domain"
	"now-go-kon/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type UserController struct {
	service application.UserService
}

func NewUserController(service application.UserService) *UserController {
	return &UserController{service: service}
}

type CreateUserRequest struct {
	Username string `json:"user_name" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type UserCreateResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user *domain.User) UserCreateResponse {
	return UserCreateResponse{
		Username: user.UserName.String(),
		Email:    user.Email.String(),
		// PasswordChangedAt: user.PasswordChangedAt.String(),
		// CreatedAt:         user.CreatedAt.String(),
	}
}

// CreateUserHandler GoDoc
// @Summary           ユーザー作成 API
// @Description       ユーザー作成時呼ばれる API
// @Param             params body CreateUserRequest true "Username, Password, Email"
// @Response          200  {object}  UserCreateResponse
// @Router            /api/v1/users [post]
func (c *UserController) CreateUserHandler(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.NewErrResponse(http.StatusInternalServerError))
		return
	}

	uParam := &domain.User{
		UserName:       domain.UserName(req.Username),
		HashedPassword: domain.HashedPassword(hashedPassword),
		Email:          domain.Email(req.Email),
	}
	user, err := c.service.CreateUser(ctx, uParam)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, domain.NewErrResponse(http.StatusForbidden))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, domain.NewErrResponse(http.StatusInternalServerError))
		return
	}

	res := newUserResponse(user)
	ctx.JSON(http.StatusOK, res)
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

// GetProfileHandler GoDoc
// @Summary           ユーザープロフィール情報参照 API
// @Description       指定ユーザーのプロフィール確認時呼ばれる API
// @Param             params body UserRequest true "ID"
// @Response          200  {object}  UserResponse
// @Router            /api/v1/users [get]
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

// UpdateProfileHandler GoDoc
// @Summary           ユーザープロフィール情報更新 API
// @Description       自身のプロフィール情報更新時呼ばれる API
// @Param             params body UserRequest true "ID, DateOfBirth, Gender, Residence, Occupation, Height, Weight"
// @Response          200  {object}  UserDetailResponse
// @Router            /api/v1/users [put]
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
