package models

type Event struct {
	Title  string `json:"title"`
	Date   string `json:"date"`
	Meta   string `json:"meta"`
	Folder string `json:"folder"`
}
