package main

import (
	"log"

	"now-go-kon/pkg"
	"now-go-kon/pkg/infrastructure"
	"now-go-kon/pkg/util"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	setupRDB(config)
	r := gin.Default()

	pkg.RegisterHandlers(r)

	r.GET("/healthCheck", pkg.HealthCheckHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setupRDB(config util.Config) {
	if err := infrastructure.RDBConnect(config); err != nil {
		log.Fatal(err)
	}
}
