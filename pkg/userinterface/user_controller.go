package userinterface

import (
	"log"
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
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	CreatedDate time.Time `json:"created_at"`
	UpdatedDate time.Time `json:"password_changed_at"`
}

func newUserResponse(user *domain.User) UserCreateResponse {
	return UserCreateResponse{
		ID:          user.ID.Num(),
		Username:    user.UserName.String(),
		Email:       user.Email.String(),
		CreatedDate: user.CreateDate,
		UpdatedDate: user.UpdatedDate,
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
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
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
	ID                 int                `json:"id"`
	Username           string             `json:"user_name"`
	Status             int                `json:"status"`
	Email              string             `json:"email"`
	UserDetailResponse UserDetailResponse `json:"user_detail"`
}

func UserProfileResponse(u *domain.User) UserResponse {
	return UserResponse{
		ID:                 u.ID.Num(),
		Username:           u.UserName.String(),
		Status:             u.Status.Num(),
		Email:              u.Email.String(),
		UserDetailResponse: UserUpdateResponse(&u.UsersDetails),
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
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Gender       string `json:"gender"`
	Height       int    `json:"height"`
	Location     string `json:"location"`
	Work         string `json:"work"`
	Graduation   string `json:"graduation"`
	Hobby        string `json:"hobby"`
	Passion      string `json:"passion"`
	Tweet        string `json:"tweet"`
	Introduction string `json:"introduction"`
}

type UserDetailResponse struct {
	UserID       int    `json:"user_id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Gender       string `json:"gender"`
	Height       int    `json:"height"`
	Location     string `json:"location"`
	Work         string `json:"work"`
	Graduation   string `json:"graduation"`
	Hobby        string `json:"hobby"`
	Passion      string `json:"passion"`
	Tweet        string `json:"tweet"`
	Introduction string `json:"introduction"`
}

func UserUpdateResponse(u *domain.UsersDetails) UserDetailResponse {
	return UserDetailResponse{
		UserID:       u.UserID.Num(),
		Name:         u.Name,
		Age:          u.Age,
		Gender:       u.Gender,
		Height:       u.Height,
		Location:     u.Location,
		Work:         u.Work,
		Graduation:   u.Graduation,
		Hobby:        u.Hobby,
		Passion:      u.Passion,
		Tweet:        u.Tweet,
		Introduction: u.Introduction,
	}
}

func (r *UserUpdateRequest) toParams() *domain.UsersDetails {
	return &domain.UsersDetails{
		UserID:       domain.UserID(r.ID),
		Name:         r.Name,
		Age:          r.Age,
		Gender:       r.Gender,
		Height:       r.Height,
		Location:     r.Location,
		Work:         r.Work,
		Graduation:   r.Graduation,
		Hobby:        r.Hobby,
		Passion:      r.Passion,
		Tweet:        r.Tweet,
		Introduction: r.Introduction,
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
	updateUser, err := c.service.UserUpsert(ctx, uParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	res := UserUpdateResponse(updateUser)
	ctx.JSON(http.StatusOK, res)
}
