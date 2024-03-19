package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/luiisp/go-uplift/src/controller/routes"
)

func main() {

	//envs := godotenv.Load()
	//if envs == nil {
	//	log.Fatal("Error in load envs")
	//}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}

}
