package models

type Event struct {
	Title  string `json:"title"`
	Date   string `json:"date"`
	Meta   string `json:"meta"`
	Status string `json:"status"`
}
