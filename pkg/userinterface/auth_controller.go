package userinterface

import (
	"net/http"
	"now-go-kon/pkg/application"
	"now-go-kon/pkg/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service application.AuthService
}

func NewAuthController(service application.AuthService) *AuthController {
	return &AuthController{service: service}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	// SessionID             uuid.UUID          `json:"session_id"`
	AccessToken           string             `json:"access_token"`
	AccessTokenExpiresAt  time.Time          `json:"access_token_expires_at"`
	RefreshToken          string             `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time          `json:"refresh_token_expires_at"`
	User                  UserCreateResponse `json:"user"`
}

// LoginHandler GoDoc
// @Summary           ログイン API
// @Description       ユーザーがログイン時呼ばれる API
// @Param             params body CreateUserRequest true "Username, Password"
// @Response          200  {object}  LoginUserResponse
// @Router            /api/v1/users/login [post]
func (c *AuthController) LoginHandler(ctx *gin.Context) {
	// var req LoginRequest
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
	// 	return
	// }

	// user, err := c.service.GetUser(ctx, req.Username)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println(err)
	// 		ctx.JSON(http.StatusNotFound, domain.NewErrResponse(http.StatusNotFound))
	// 		return
	// 	}
	// 	log.Println(err)
	// 	ctx.JSON(http.StatusInternalServerError, domain.NewErrResponse(http.StatusInternalServerError))
	// 	return
	// }
}

// LoginHandler GoDoc
// @Summary           セッション確認API API
// @Description       セッション確認時呼ばれる API
// @Param             params body CreateUserRequest true ""
// @Response          200  {object}  LoginUserResponse
// @Router            /api/v1/users/login [post]
func (c *AuthController) GetSessionHandler(ctx *gin.Context) {

}

// LoginHandler GoDoc
// @Summary           ログアウト API
// @Description       ユーザーがログアウト時呼ばれる API
// @Param             params body CreateUserRequest true ""
// @Response          200  {object}  LoginUserResponse
// @Router            /api/v1/users/login [post]
func (c *AuthController) LogoutHandler(ctx *gin.Context) {

}

type AuthRequest struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

// LoginHandler GoDoc
// @Summary           パスワード認証 API
// @Description       ユーザーがログイン時呼ばれる API
// @Param             params body CreateUserRequest true "Username, Password"
// @Response          200  {object}  LoginUserResponse
// @Router            /api/v1/users/login [post]
func (c *AuthController) PasswordAuthHandler(ctx *gin.Context) {
	var req AuthRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	if err := c.service.Auth(ctx, domain.UserID(req.ID), domain.Password(req.Password)); err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}
	ctx.Status(http.StatusOK)
}
