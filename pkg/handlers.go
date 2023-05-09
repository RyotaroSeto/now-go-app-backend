package pkg

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"now-go-kon/pkg/injection"
	"now-go-kon/pkg/token"
	"now-go-kon/pkg/util"
)

type Health struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func HealthCheckHandler(c *gin.Context) {
	health := Health{
		Status: http.StatusOK,
		Result: "success",
	}
	c.JSON(200, health)
}

func RegisterHandlers(e *gin.Engine, config util.Config) {
	root := e.Group("/api/v1")

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker: %w", err)
	}

	{
		RegisterAuthenticationHandlers(root, tokenMaker, config)
		RegisterUserHandlers(root, tokenMaker)
		RegisterBoardHandlers(root, tokenMaker)
		RegisterLikeHandlers(root, tokenMaker)
	}
}

func RegisterUserHandlers(root *gin.RouterGroup, token token.Maker) {
	user := injection.InitializeUserController()

	authRoutes := root.Group("/users").Use(authMiddleware(token))
	users := root.Group("/users")
	{
		users.POST("/", user.CreateUserHandler)
		authRoutes.GET("/", user.GetProfileHandler)
		authRoutes.POST("/upsert", user.UpdateProfileHandler)
	}
}

func RegisterAuthenticationHandlers(root *gin.RouterGroup, token token.Maker, config util.Config) {
	auth := injection.InitializeAuthController()

	authRoutes := root.Group("/session").Use(authMiddleware(token))
	session := root.Group("/session")
	{
		session.POST("/login", func(ctx *gin.Context) {
			auth.LoginHandler(ctx, token, config)
		})
		authRoutes.GET("/", auth.GetSessionHandler)
		authRoutes.DELETE("/", auth.LogoutHandler)
		authRoutes.POST("/", auth.PasswordAuthHandler)
	}
}

func RegisterBoardHandlers(root *gin.RouterGroup, token token.Maker) {
	board := injection.InitializeBoardController()

	authRoutes := root.Group("/board").Use(authMiddleware(token))
	{
		authRoutes.GET("/", board.GetBoardHandler)
		authRoutes.GET("/scroll", board.GetBoardScrollHandler)
		authRoutes.POST("/", board.CreateBoardHandler)
		authRoutes.DELETE("/", board.DeleteBoardHandler)
	}
}

func RegisterLikeHandlers(root *gin.RouterGroup, token token.Maker) {
	like := injection.InitializeLikeController()

	authRoutes := root.Group("/like").Use(authMiddleware(token))
	{
		authRoutes.GET("/", like.GetLikeHandler)
		authRoutes.POST("/", like.CreateLikeHandler)
		authRoutes.POST("/approval", like.ApprovalHandler)

	}
}

func RegisterMessageHandlers(root *gin.RouterGroup, token token.Maker) {
	message := injection.InitializeMessageController()

	authRoutes := root.Group("/message").Use(authMiddleware(token))
	{
		authRoutes.GET("/", message.GetMessageHandler)
		authRoutes.GET("/scroll", message.GetMessageScrollHandler)
		authRoutes.POST("/", message.CreateMessageHandler)
	}
}

func RegisterNotFoundHandler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
		err := fmt.Errorf(
			"<method: %s, url: %s, params: %+v> is not found in routes",
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.URL.Query(),
		)
		log.Println(err)
	})
}
