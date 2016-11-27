package handlers

import (
	// "encoding/json"
	"net/http"
	// log "github.com/Sirupsen/logrus"
)

type Components struct {
	*application
}

func (p *Components) List(req *http.Request) ([]byte, *HttpError) {

	dataStr := `
	{
		"data": [
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
	
		],
		"message": "成功",
		"state_code": 200	
	}
	`

	return []byte(dataStr), nil
}
