package models

type Concept struct {
	Uuid          string   `json:"uuid"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	Prerequisites []string `json:"prerequisites"`
}
