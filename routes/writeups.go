package routes

import (
	"github.com/gin-gonic/gin"
	"hofill/api_status"
	"hofill/models"
	"hofill/repository"
	"net/http"
)

type WriteUpHandler struct {
	repo repository.WriteUpRepository
}

func (w WriteUpHandler) getWriteUpsForEvent(c *gin.Context) {
	if eventName := c.Query("event-name"); eventName != "" {
		c.IndentedJSON(http.StatusOK, w.repo.GetWriteUpPreviewsForEvent(eventName))
		return
	}
	c.IndentedJSON(http.StatusUnprocessableEntity, models.WriteUpPreview{Status: api_status.NotFoundError})
}

func SetupWriteUpRoutes(r *gin.Engine, repo repository.WriteUpRepository) {
	handler := WriteUpHandler{repo}
	r.GET("/writeups", handler.getWriteUpsForEvent)
	r.GET("/writeup")
}
