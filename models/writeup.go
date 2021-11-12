package models

type WriteUp struct {
	MDData         string `json:"mdData" mapstructure:"md_data"`
	WriteUpPreview `mapstructure:",squash"`
}

type WriteUpPreview struct {
	Title      string `json:"title"`
	Date       string `json:"date"`
	Event      string `json:"event"`
	Category   string `json:"category"`
	Difficulty string `json:"difficulty"`
	FileName   string `json:"fileName" mapstructure:"file_name"`
	Points     int    `json:"points"`
	Solved     bool   `json:"solved"`
	Flag       string `json:"flag"`
}
