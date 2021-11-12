package routes

import (
	"github.com/gin-gonic/gin"
	"hofill/models"
	"hofill/repository"
	"net/http"
)

type WriteUpHandler struct {
	repo repository.WriteUpRepository
}

func (w WriteUpHandler) getWriteUpsForEvent(c *gin.Context) {
	var status = http.StatusNotFound
	var writeUpPreviews []models.WriteUpPreview
	if eventName := c.Query("event-name"); eventName != "" {
		writeUpPreviews := w.repo.GetWriteUpPreviewsForEvent(eventName)
		if len(writeUpPreviews) != 0 {
			status = http.StatusOK
		}
	}
	c.IndentedJSON(status, writeUpPreviews)
}

func (w WriteUpHandler) getWriteUp(c *gin.Context) {
	eventName, writeUpName := c.Query("event-name"), c.Query("writeup-name")
	var status = http.StatusNotFound
	var writeUp models.WriteUp
	var err error
	if eventName != "" && writeUpName != "" {
		writeUp, err = w.repo.GetWriteUp(eventName, writeUpName)
		if err == nil {
			status = http.StatusOK
		}
	}
	c.IndentedJSON(status, writeUp)
}

func SetupWriteUpRoutes(r *gin.Engine, repo repository.WriteUpRepository) {
	handler := WriteUpHandler{repo}
	r.GET("/writeups", handler.getWriteUpsForEvent)
	r.GET("/writeup", handler.getWriteUp)
}
