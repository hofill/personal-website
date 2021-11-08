package models

type WriteUp struct {
}

type WriteUpPreview struct {
	Title    string `json:"title"`
	Date     string `json:"date"`
	Event    string `json:"event"`
	Category string `json:"category"`
}
