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
				"icon": "http://ohan0t6mr.bkt.clouddn.com/assets/components/huodong.png",
				"api_path": "` + activityPath + "&page_no=3" + `"
			},			
			{
				"name": "最热",
				"icon": "http://ohan0t6mr.bkt.clouddn.com/assets/components/zuire.png",
				"api_path": "` + activityPath + "&page_no=2" + `"
			},
			
			{
				"name": "推荐",
				"icon": "http://ohan0t6mr.bkt.clouddn.com/assets/components/tuijian.png",
				"api_path": "` + activityPath + "&page_no=4" + `"
			},
			
			{
				"name": "破解",
				"icon": "http://ohan0t6mr.bkt.clouddn.com/assets/components/pojie.png",
				"api_path": "` + activityPath + "&page_no=5" + `"
			}
	
		],
		"message": "成功",
		"state_code": 200	
	}
	`

	return []byte(dataStr), nil
}
