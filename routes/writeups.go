package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"hofill/api_status"
	"hofill/models"
	"hofill/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getWriteUpsForEvent(c *gin.Context) {
	if eventName := c.Query("event-name"); eventName != "" {
		c.IndentedJSON(http.StatusOK, getAllWriteUps(eventName))
		return
	}
	c.IndentedJSON(http.StatusUnprocessableEntity, models.WriteUpPreview{Status: api_status.NotFoundError})

}

func SetupWriteUpRoutes(r *gin.Engine) {
	r.GET("/writeups", getWriteUpsForEvent)
	r.GET("/writeup")
}

func getAllWriteUps(eventName string) []models.WriteUpPreview {
	const eventsFolderPath = "./_events"
	var selectedEventFolderPath = eventsFolderPath + "/" + eventName + "/"
	var writeUps []models.WriteUpPreview

	// Check if given event folder exists
	files, err := ioutil.ReadDir(selectedEventFolderPath)
	if err != nil {
		log.Println(err)
		return append(writeUps, models.WriteUpPreview{Status: api_status.FileError})
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") || !strings.HasSuffix(file.Name(), ".md") {
			continue
		}

		writeUpsMap, err := utils.GetWriteUpMetaData(selectedEventFolderPath + file.Name())
		if err != nil {
			log.Println(err)
			continue
		}

		var writeUp models.WriteUpPreview
		err = mapstructure.Decode(writeUpsMap, &writeUp)
		if err != nil {
			log.Println(err)
			continue
		}

		writeUps = append(writeUps, writeUp)
	}

	if writeUps == nil {
		return append(writeUps, models.WriteUpPreview{Status: api_status.NoWriteUpFoundError})
	}

	return writeUps

}
