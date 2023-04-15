package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	setupRDB()
	r := gin.Default()

	RegisterHandlers(r)

	r.GET("/healthCheck", HealthCheckHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setupRDB() {
	// authDSN := "認証コンテキスト用DBへの接続情報"
	// if err := authdb.RDBConnect(authDSN); err != nil {
	// 	log.Fatal(err)
	// }
}
