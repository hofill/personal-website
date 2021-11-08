package main

import (
	"github.com/gin-gonic/gin"
	"hofill/routes"
)

func main() {
	startAPI()
}

func startAPI() {
	router := gin.Default()
	router.GET("/events", routes.GetEvents)

	err := router.Run("localhost:443")
	if err != nil {
		return 
	}
}
