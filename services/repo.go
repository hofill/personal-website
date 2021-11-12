package services

import (
	"github.com/mitchellh/mapstructure"
	"hofill/models"
	"hofill/repository"
	"hofill/utils"
	"io/ioutil"
	"log"
	"strings"
)

const (
	eventFolderPath  = "./_events"
	metaDataFileName = ".metadata"
	writeUpExtension = ".md"
)

func PopulateRepository(repo repository.WriteUpRepository) {
	populateEvents(repo)
	populateWriteUps(repo)
}

func populateEvents(repo repository.WriteUpRepository) {

	files, err := ioutil.ReadDir(eventFolderPath)
	if err != nil {
		log.Fatal(err)
		return
	}

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

		event.Folder = f.Name()
		repo.AddEvent(event)
	}

}

func populateWriteUps(repo repository.WriteUpRepository) {
	for _, event := range repo.GetEvents() {
		var selectedEventFolderPath = eventFolderPath + "/" + event.Folder + "/"

		files, err := ioutil.ReadDir(selectedEventFolderPath)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, file := range files {
			if strings.HasPrefix(file.Name(), ".") || !strings.HasSuffix(file.Name(), writeUpExtension) {
				continue
			}

			writeUpMap, writeUpMD, err := utils.GetWriteUpMetaDataAndMD(selectedEventFolderPath + file.Name())
			if err != nil {
				log.Println(err)
				continue
			}

			var writeUp models.WriteUp
			err = mapstructure.Decode(writeUpMap, &writeUp)
			if err != nil {
				log.Println(err)
				continue
			}
			writeUp.MDData = writeUpMD

			repo.AddWriteUp(writeUp)
		}
	}
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
