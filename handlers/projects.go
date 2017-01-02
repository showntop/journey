package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"

	"github.com/showntop/journey/models"
	. "github.com/showntop/journey/stores"
)

type Projects struct {
	*application
}

func (p *Projects) List(req *http.Request) ([]byte, *HttpError) {
	queryValues := req.URL.Query()

	//conditions
	var categoryId, tagId int64
	_ = tagId
	var pageNo, pageNum int
	// var key string
	var err error

	if pno := queryValues.Get("page_no"); pno != "" {
		if pageNo, err = strconv.Atoi(pno); err != nil {
			pageNo = 1 //default
		}
	}
	if pnum := queryValues.Get("page_num"); pnum != "" {
		if pageNum, err = strconv.Atoi(pnum); err != nil {
			pageNum = 15 //default
		}
	}

	// var projects []*models.Project
	// if categoryId <= 0 {
	// 	projects, err = StoreM.Project.FindAll(pageNo, pageNum)
	// } else {
	// 	projects, err = StoreM.Project.FindAllByCategory(categoryId, pageNo, pageNum)
	// }

	categoryId, _ = strconv.ParseInt(req.URL.Query().Get("category_id"), 10, 64)

	var projects []*models.Project
	if categoryId <= 0 {
		projects, err = StoreM.Project.FindAll((pageNo-1)*pageNum, pageNum)
	} else {
		projects, err = StoreM.Project.FindAllByCategory(categoryId, (pageNo-1)*pageNum, pageNum)
	}
	if err != nil {
		log.Errorln("projects database error", err)
		return nil, DBErr
	}
	output, err := json.Marshal(WrapRespData(projects))
	if err != nil {
		log.Warnln("projects marshal projects,", err)
		return output, BadRespErr
	}
	return output, nil
}

func (p *Projects) Show(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	pids := ps.ByName("id")
	pidi, err := strconv.ParseInt(pids, 10, 64)
	if err != nil {
		log.Errorln("projects database error", err)
		return nil, BadRequestErr
	}

	project, err := StoreM.Project.Find(pidi)
	if err != nil {
		log.Errorln("projects database error", err)
		return nil, DBErr
	}

	output, err := json.Marshal(WrapRespData(project))
	if err != nil {
		log.Warnln("projects marshal projects,", err)
		return output, BadRespErr
	}
	return output, nil
}

func (p *Projects) Search(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	query := req.URL.Query()
	key := query.Get("key")

	projects, err := StoreM.Project.FindWithKey(key, 0, 10)
	if err != nil {
		return nil, DBErr
	}
	output, err := json.Marshal(WrapRespData(projects))
	if err != nil {
		log.Warnln("projects marshal projects,", err)
		return output, BadRespErr
	}
	return output, nil
}
