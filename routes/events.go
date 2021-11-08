package routes

import (
	"github.com/gin-gonic/gin"
	"hofill/models"
	"io/ioutil"
	"log"
	"net/http"
)

func GetEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getEventFolders())
}

func getEventFolders() []models.Event {
	files, err := ioutil.ReadDir("./events")
	if err != nil {
		log.Fatal(err)
	}

	var events []models.Event

	for _, f := range files {
		var event = models.Event{
			Title: f.Name(),
			Date:  f.ModTime().String(),
		}
		events = append(events, event)
	}

	return events
}
