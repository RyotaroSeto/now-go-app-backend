package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func main() {
	setupRDB()
	r := gin.Default()

	RegisterHandlers(r)

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"response": "test"})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setupRDB() {
	// authDSN := "認証コンテキスト用DBへの接続情報"
	// if err := authdb.RDBConnect(authDSN); err != nil {
	// 	log.Fatal(err)
	// }

	// userDSN := "ユーザー管理コンテキスト用DBへの接続情報"
	// if err := userdb.RDBConnect(userDSN); err != nil {
	// 	log.Fatal(err)
	// }
}
