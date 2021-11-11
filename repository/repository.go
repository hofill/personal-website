package repository

import "hofill/models"

type WriteUpRepository interface {
	addEvent(event models.Event)

	removeEvent(event models.Event) error

	addWriteUp(writeUp models.WriteUp)

	removeWriteUp(writeUp models.WriteUp) error

	getEvents() []models.Event

	orderEventsByDate() []models.Event

	getWriteUps() []models.WriteUp

	getWriteUpPreviews() []models.WriteUpPreview

	getWriteUpsForEvent(eventName string) []models.WriteUp

	getWriteUpPreviewsForEvent(eventName string) []models.WriteUpPreview
}
