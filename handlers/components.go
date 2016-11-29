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
				"name": "活动",
				"icon": "http://ohan0t6mr.bkt.clouddn.com/activity.png",
				"api_path": ""
			},			
			{
				"name": "最热",
				"icon": "http://ohan0t6mr.bkt.clouddn.com/hottest.png",
				"api_path": ""
			},
			
			{
				"name": "最新",
				"icon": "http://ohan0t6mr.bkt.clouddn.com/latest.png",
				"api_path": ""
			},
			
			{
				"name": "推荐",
				"icon": "",
				"api_path": "http://ohan0t6mr.bkt.clouddn.com/recommend.png"
			},
			
			{
				"name": "破解",
				"icon": "http://ohan0t6mr.bkt.clouddn.com/unlock.png",
				"api_path": ""
			}
	
		],
		"message": "成功",
		"state_code": 200	
	}
	`

	return []byte(dataStr), nil
}
