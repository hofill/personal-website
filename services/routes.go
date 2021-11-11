package services

import (
	"github.com/gin-gonic/gin"
	"hofill/repository"
	"hofill/routes"
	"log"
	"os"
)

func SetupRoutes(repo repository.WriteUpRepository) {
	router := gin.Default()

	routes.SetupEventRoutes(router, repo)
	routes.SetupWriteUpRoutes(router, repo)
	routes.SetupDefaultRoutes(router)

	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

}
