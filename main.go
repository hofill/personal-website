package main

import (
	"github.com/gin-gonic/gin"
	"hofill/routes"
	"log"
	"os"
)

func main() {
	startAPI()
}

func startAPI() {
	router := gin.Default()

	routes.SetupEventRoutes(router)
	routes.SetupWriteUpRoutes(router)
	routes.SetupDefaultRoutes(router)

	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
