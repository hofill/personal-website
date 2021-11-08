package main

import (
	"github.com/gin-gonic/gin"
	"hofill/routes"
	"net/http"
	"os"
)

func main() {
	startAPI()
}

func startAPI() {
	router := gin.Default()
	router.GET("/events", routes.GetEvents)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "server running")
	})

	err := router.Run("localhost:" + os.Getenv("PORT"))
	if err != nil {
		return 
	}
}
