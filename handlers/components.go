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
				"icon": "http://oizfueie4.bkt.clouddn.com/assets/components/activity.png",
				"api_path": "http://192.168.1.113:8000/api/v1/apps?page_num=10"
			},			
			{
				"name": "最热",
				"icon": "http://oizfueie4.bkt.clouddn.com/assets/components/hot.png",
				"api_path": "http://192.168.1.113:8000/api/v1/apps?page_num=10"
			},
			
			{
				"name": "推荐",
				"icon": "http://oizfueie4.bkt.clouddn.com/assets/components/recommend.png",
				"api_path": "http://192.168.1.113:8000/api/v1/apps?page_num=10"
			},
			
			{
				"name": "破解",
				"icon": "http://oizfueie4.bkt.clouddn.com/assets/components/unlock.png",
				"api_path": "http://192.168.1.113:8000/api/v1/apps?page_num=10"
			}
	
		],
		"message": "成功",
		"state_code": 200	
	}
	`

	return []byte(dataStr), nil
}
