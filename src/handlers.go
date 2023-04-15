package main

import (
	"github.com/gin-gonic/gin"

	authInjection "authentication/injection"
	// userInjection "user/injection"
)

func RegisterHandlers(e *gin.Engine) {
	root := e.Group("/api/v1")

	{
		RegisterAuthenticationHandlers(root)
		// RegisterContractHandlers(root)
	}
}

// func RegisterUserHandlers(root *gin.RouterGroup) {
// 	user := userInjection.InitializeUser()

// 	users := root.Group("/users")
// 	{
// 		users.GET("/profile", user.GetProfileHandler) //ユーザー情報参照API。GET /api/v1/users/profile
// 	}
// }

func RegisterAuthenticationHandlers(root *gin.RouterGroup) {
	auth := authInjection.InitializeAuthController()

	session := root.Group("/session")
	{
		// session.GET("/", auth.GetSessionHandler)    //セッション確認API。GET /api/v1/session/
		session.POST("/", auth.PasswordAuthHandler) //パスワード認証API。POST /api/v1/session/
		// session.DELETE("/", auth.LogoutHandler)     //ログアウトAPI。DELETE /api/v1/session/
	}
}
