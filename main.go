package main

import (
	"github.com/gin-gonic/gin"
	"hofill/routes"
	"net/http"
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

	err := router.Run("localhost:8080")
	if err != nil {
		return 
	}
}
