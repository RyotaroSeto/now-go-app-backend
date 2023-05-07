package userinterface

import (
	"database/sql"
	"log"
	"net/http"
	"now-go-kon/pkg/application"
	"now-go-kon/pkg/domain"
	"now-go-kon/pkg/token"
	"now-go-kon/pkg/util"
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
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func toParams(refreshPayload *token.Payload, username string, refreshToken string, ctx *gin.Context) *domain.Session {
	return &domain.Session{
		SessionID:    domain.SessionID(refreshPayload.ID),
		UserName:     domain.UserName(username),
		RefreshToken: domain.RefreshToken(refreshToken),
		UserAgent:    domain.UserAgent(ctx.Request.UserAgent()),
		ClientIP:     domain.ClientIP(ctx.ClientIP()),
		IsBlocked:    false,
		ExpiresDate:  refreshPayload.ExpiredAt,
	}
}

type loginUserResponse struct {
	SessionID             string             `json:"session_id"`
	AccessToken           string             `json:"access_token"`
	AccessTokenExpiresAt  time.Time          `json:"access_token_expires_at"`
	RefreshToken          string             `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time          `json:"refresh_token_expires_at"`
	User                  UserCreateResponse `json:"user"`
}

// LoginHandler GoDoc
// @Summary           ログイン API
// @Description       ユーザーがログイン時呼ばれる API
// @Param             params body CreateUserRequest true "Email, Password"
// @Response          200  {object}  loginUserResponse
// @Router            /api/v1/users/login [post]
func (c *AuthController) LoginHandler(ctx *gin.Context, tokenMaker token.Maker, config util.Config) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewErrResponse(http.StatusBadRequest))
		return
	}

	email := domain.Email(req.Email)
	user, err := c.service.GetUser(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			ctx.JSON(http.StatusNotFound, domain.NewErrResponse(http.StatusNotFound))
			return
		}
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, domain.NewErrResponse(http.StatusInternalServerError))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword.String())
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusUnauthorized, domain.NewErrResponse(http.StatusUnauthorized))
		return
	}

	accessToken, accessPayload, err := tokenMaker.CreateToken(
		user.UserName.String(),
		config.AccessTokenDuration,
	)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, domain.NewErrResponse(http.StatusInternalServerError))
		return
	}

	refreshToken, refreshPayload, err := tokenMaker.CreateToken(
		user.UserName.String(),
		config.RefreshTokenDuration,
	)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, domain.NewErrResponse(http.StatusInternalServerError))
		return
	}

	sParam := toParams(refreshPayload, user.UserName.String(), refreshToken, ctx)
	session, err := c.service.CreateSession(ctx, sParam)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, domain.NewErrResponse(http.StatusInternalServerError))
		return
	}
	log.Println(111111111)
	log.Println(session)

	rsp := loginUserResponse{
		// SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)

}

func (c *AuthController) GetSessionHandler(ctx *gin.Context) {

}

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
