package models

type ConceptNode struct {
	Uuid    string `json:"uuid"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ConceptLink struct {
	StartId string `json:"start_id"`
	EndId   string `json:"end_id"`
}
