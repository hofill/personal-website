package repository

import (
	"errors"
	"github.com/jinzhu/copier"
	"hofill/models"
	"hofill/utils"
	"sort"
)

type InMemoryRepository struct {
	events   []models.Event
	writeUps []models.WriteUp
}

func (i *InMemoryRepository) AddEvent(event models.Event) {
	i.events = append(i.events, event)
}

func (i *InMemoryRepository) RemoveEvent(event models.Event) error {
	for index, e := range i.events {
		if e == event {
			i.events = append(i.events[:index], i.events[index+1:]...)
			return nil
		}
	}
	return errors.New("event not found")
}

func (i *InMemoryRepository) AddWriteUp(writeUp models.WriteUp) {
	i.writeUps = append(i.writeUps, writeUp)
}

func (i *InMemoryRepository) RemoveWriteUp(writeUp models.WriteUp) error {
	for index, w := range i.writeUps {
		if w == writeUp {
			i.writeUps = append(i.writeUps[:index], i.writeUps[index+1:]...)
			return nil
		}
	}
	return errors.New("writeup not found")
}

func (i *InMemoryRepository) GetEvents() []models.Event {
	return i.events
}

func (i *InMemoryRepository) OrderEventsByDate() []models.Event {
	var events = i.events
	sort.Slice(events, func(i, j int) bool {
		timeFirst := utils.TimeToLayoutUS(events[i].Date)
		timeSecond := utils.TimeToLayoutUS(events[j].Date)
		return timeFirst.After(timeSecond)
	})
	return events
}

func (i *InMemoryRepository) GetWriteUps() []models.WriteUp {
	return i.writeUps
}

func (i *InMemoryRepository) GetWriteUpPreviews() []models.WriteUpPreview {
	var writeUpPreviews []models.WriteUpPreview

	for _, writeUp := range i.writeUps {
		var writeUpPreview models.WriteUpPreview
		err := copier.Copy(&writeUpPreview, writeUp)
		if err != nil {
			continue
		}
		writeUpPreviews = append(writeUpPreviews, writeUpPreview)
	}

	return writeUpPreviews
}

func (i *InMemoryRepository) GetWriteUpsForEvent(eventName string) []models.WriteUp {
	var writeUpsForEvent []models.WriteUp
	for _, writeUp := range i.writeUps {
		if writeUp.Event == eventName {
			writeUpsForEvent = append(writeUpsForEvent, writeUp)
		}
	}
	return writeUpsForEvent
}

func (i *InMemoryRepository) GetWriteUpPreviewsForEvent(eventName string) []models.WriteUpPreview {
	var writeUpPreviewsForEvent []models.WriteUpPreview
	for _, writeUp := range i.GetWriteUpPreviews() {
		if writeUp.Event == eventName {
			writeUpPreviewsForEvent = append(writeUpPreviewsForEvent, writeUp)
		}
	}
	return writeUpPreviewsForEvent
}
