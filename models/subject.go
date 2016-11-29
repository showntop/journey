package models

type Subject struct {
	Base
	Title   string `json:"title"`
	Content string `json:"content"`
	Asset   string `json:"asset"`
	Action  string `json:"action"`
	Params  string `json:"params"`
}
