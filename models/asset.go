package models

type Asset struct {
	Base
	Type   int    `json:"-"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Action string `json:"action"`
	Params string `json:"params"`
}
