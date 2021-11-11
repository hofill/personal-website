package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mitchellh/mapstructure"
	"hofill/repository"
	"net/http"
)

type EventHandler struct {
	repo repository.WriteUpRepository
}

func (e EventHandler) getEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, e.repo.OrderEventsByDate())
}

func SetupEventRoutes(r *gin.Engine, repo repository.WriteUpRepository) {
	eventHandler := EventHandler{repo}
	r.GET("/events", eventHandler.getEvents)
}
