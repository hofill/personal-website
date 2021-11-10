package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	_ "github.com/mitchellh/mapstructure"
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
		event, err := getEventData(eventMetaDataPath)
		if err != nil {
			log.Println(err)
			continue
		}

		if event.Title == "" {
			event.Title = f.Name()
		}

		event.Status = "OK"
		events = append(events, event)
	}

	return events
}

func getEventData(fileName string) (models.Event, error) {
	event := models.Event{}

	metaDataValues, err := utils.GetEventMetaData(fileName)
	if err != nil {
		return models.Event{}, err
	}

	err = mapstructure.Decode(metaDataValues, &event)
	if err != nil {
		return event, err
	}

	return event, nil

}
