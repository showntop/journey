package models

type Asset struct {
	Base
	Name string `json:"name"`
	URL  string `json:"url"`
}
