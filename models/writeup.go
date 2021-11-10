package models

type WriteUp struct {
	MDData string `json:"md_data"`
	WriteUpPreview
}

type WriteUpPreview struct {
	Title      string `json:"title"`
	Date       string `json:"date"`
	Event      string `json:"event"`
	Category   string `json:"category"`
	Difficulty string `json:"difficulty"`
	Points     string `json:"points"`
	Solved     bool   `json:"solved"`
	Flag       string `json:"flag"`
	Status     string `json:"status"`
}
