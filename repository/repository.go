package repository

import "hofill/models"

type WriteUpRepository interface {
	AddEvent(event models.Event)

	RemoveEvent(event models.Event) error

	AddWriteUp(writeUp models.WriteUp)

	RemoveWriteUp(writeUp models.WriteUp) error

	GetEvents() []models.Event

	OrderEventsByDate() []models.Event

	GetWriteUps() []models.WriteUp

	GetWriteUp(eventName, writeUpName string) (models.WriteUp, error)

	GetWriteUpPreviews() []models.WriteUpPreview

	GetWriteUpsForEvent(eventName string) []models.WriteUp

	GetWriteUpPreviewsForEvent(eventName string) []models.WriteUpPreview
}
