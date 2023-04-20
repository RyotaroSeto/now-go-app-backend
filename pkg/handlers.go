package pkg

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"now-go-kon/pkg/injection"
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

func RegisterHandlers(e *gin.Engine) {
	root := e.Group("/api/v1")

	{
		RegisterAuthenticationHandlers(root)
		RegisterUserHandlers(root)
	}
}

func RegisterUserHandlers(root *gin.RouterGroup) {
	user := injection.InitializeUserController()

	users := root.Group("/users")
	{
		users.POST("/profile", user.GetProfileHandler)
		users.PUT("/update", user.UpdateProfileHandler)
	}
}

func RegisterAuthenticationHandlers(root *gin.RouterGroup) {
	auth := injection.InitializeAuthController()

	session := root.Group("/session")
	{
		// session.GET("/", auth.GetSessionHandler)    //セッション確認API。GET /api/v1/session/
		session.POST("/", auth.PasswordAuthHandler) //パスワード認証API。POST /api/v1/session/
		// session.DELETE("/", auth.LogoutHandler)     //ログアウトAPI。DELETE /api/v1/session/
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
