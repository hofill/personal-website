package services

import (
	"github.com/gin-gonic/gin"
	"hofill/routes"
	"log"
	"os"
)

func SetupRoutes() {
	router := gin.Default()

	routes.SetupEventRoutes(router)
	routes.SetupWriteUpRoutes(router)
	routes.SetupDefaultRoutes(router)

	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

}
