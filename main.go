package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luiisp/go-uplift/src/configuration/logger"
	"github.com/luiisp/go-uplift/src/controller/routes"
)

func main() {
	logger.Info("Starting Application")
	envs := godotenv.Load(".env")
	if envs == nil {
		logger.Error("Error in load envs", nil)
	}
	logger.Info("Loaded envs.")

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}
