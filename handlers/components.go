package handlers

import (
	"net/http"
)

type Components struct {
	*application
}

func (p *Components) List(req *http.Request) ([]byte, *HttpError) {
	return []byte(`
		[
			{
				"name": "最新",
				"icon": "",
				"api_path": ""
			},			
			{
				"name": "最新",
				"icon": "",
				"api_path": ""
			},
			
			{
				"name": "最新",
				"icon": "",
				"api_path": ""
			},
			
			{
				"name": "最新",
				"icon": "",
				"api_path": ""
			}
	
		]
	`), nil
}
