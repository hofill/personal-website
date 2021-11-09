package routes

import (
	"github.com/gin-gonic/gin"
	"hofill/models"
	"hofill/utils"
	"io/ioutil"
	"log"
	"net/http"
)

func getEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getEventFolders())
}

func SetupEventRoutes(r *gin.Engine) {
	r.GET("/events", getEvents)
}

func getEventFolders() []models.Event {
	const eventFolderPath = "./_events"
	const metaDataFileName = ".metadata"

	files, err := ioutil.ReadDir(eventFolderPath)
	if err != nil {
		log.Fatal(err)
	}

	var events []models.Event

	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		eventMetaDataPath := eventFolderPath +
			"/" + f.Name() + "/" +
			metaDataFileName
		var event = models.Event{
			Title:  f.Name(),
			Date:   utils.GetDateFormatted(f.ModTime()),
			Meta:   getMetaData(eventMetaDataPath),
			Status: "OK",
		}
		events = append(events, event)
	}

	return events
}

func getMetaData(fileName string) string {
	metaDataValues, err := utils.GetEventMetaData(fileName)
	if err != nil {
		return ""
	}
	if value, ok := metaDataValues["meta"]; ok {
		return value
	}
	return ""
}
