package routes

import (
	"bufio"
	"errors"
	"github.com/gin-gonic/gin"
	"hofill/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getEventFolders())
}

func SetupEventRoutes(r *gin.Engine) {
	r.GET("/events", getEvents)
}

func getEventFolders() []models.Event {
	files, err := ioutil.ReadDir("./events")
	if err != nil {
		log.Fatal(err)
	}

	var events []models.Event

	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		var event = models.Event{
			Title:  f.Name(),
			Date:   getDateFormatted(f.ModTime()),
			Meta:   getMetaData(f.Name()),
			Status: "OK",
		}
		events = append(events, event)
	}

	return events
}

func getDateFormatted(t time.Time) string {
	// format time by us layout, january 2nd 2006
	return t.Format("January 2, 2006")
}

func getMetaData(fileName string) string {
	metaDataValues, err := getKeysFromMetaDataFile(fileName)
	if err != nil {
		return ""
	}
	if value, ok := metaDataValues["meta"]; ok {
		return value
	}
	return ""
}

func getKeysFromMetaDataFile(fileName string) (map[string]string, error) {
	var filePath = "./events/" + fileName + "/.metadata"
	// Open the file and create a scanner
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var metaData = make(map[string]string)

	for scanner.Scan() {
		keyValuePair := scanner.Text()
		splitPair := strings.Split(keyValuePair, "=")
		if len(splitPair) != 2 {
			return nil, errors.New("corrupted metadata file")
		}
		key, value := splitPair[0], splitPair[1]
		metaData[key] = value
	}

	if metaData == nil {
		return nil, errors.New("corrupted metadata file")
	}

	return metaData, nil

}
