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
	r := gin.Default()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"response": "test"})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
