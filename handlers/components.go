package handlers

import (
	// "encoding/json"
	"fmt"
	"net/http"
	// log "github.com/Sirupsen/logrus"

	. "github.com/showntop/journey/config"
)

type Components struct {
	*application
}

func (p *Components) List(req *http.Request) ([]byte, *HttpError) {
	activityPath := fmt.Sprintf("%s%s", Config.Host["domain"], "/api/v1/apps?page_num=10")
	dataStr := `
	{
		"data": [
			{
				"name": "活动",
				"icon": "http://oizfueie4.bkt.clouddn.com/assets/components/activity.png",
				"api_path": "` + activityPath + "&page_no=3" + `"
			},			
			{
				"name": "最热",
				"icon": "http://oizfueie4.bkt.clouddn.com/assets/components/hot.png",
				"api_path": "` + activityPath + "&page_no=2" + `"
			},
			
			{
				"name": "推荐",
				"icon": "http://oizfueie4.bkt.clouddn.com/assets/components/recommend.png",
				"api_path": "` + activityPath + "&page_no=4" + `"
			},
			
			{
				"name": "破解",
				"icon": "http://oizfueie4.bkt.clouddn.com/assets/components/unlock.png",
				"api_path": "` + activityPath + "&page_no=5" + `"
			}
	
		],
		"message": "成功",
		"state_code": 200	
	}
	`

	return []byte(dataStr), nil
}
