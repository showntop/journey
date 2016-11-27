package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"

	// "github.com/showntop/journey/models"
	. "github.com/showntop/journey/stores"
)

type Projects struct {
	*application
}

func (p *Projects) List(req *http.Request) ([]byte, *HttpError) {
	queryValues := req.URL.Query()

	//conditions
	// var categoryId, tagId int64
	var pageNo, pageNum int
	// var key string
	var err error

	var querySQL string = "true"

	if cidstr := queryValues.Get("category_id"); cidstr != "" {
		// categoryId, _ = strconv.ParseInt(cidstr, 10, 64)
		querySQL += " and category_id = " + cidstr
	}
	if tidstr := queryValues.Get("tag_id"); tidstr != "" {
		// tagId, _ = strconv.ParseInt(tidstr, 10, 64)
		querySQL += " and tag_id = " + tidstr
	}
	log.Debugf("key:  %v", queryValues.Get("key"))
	if pkey := queryValues.Get("key"); pkey != "" {
		// key = pkey
		querySQL += " and (name like '%" + pkey + "%' or " + " description like '%" + pkey + "%')"
	}
	if pno := queryValues.Get("page_no"); pno != "" {
		if pageNo, err = strconv.Atoi(pno); err != nil {
			pageNo = 1 //default
		}
	}
	if pnum := queryValues.Get("page_num"); pnum != "" {
		if pageNo, err = strconv.Atoi(pnum); err != nil {
			pageNo = 15 //default
		}
	}

	// var projects []*models.Project
	// if categoryId <= 0 {
	// 	projects, err = StoreM.Project.FindAll(pageNo, pageNum)
	// } else {
	// 	projects, err = StoreM.Project.FindAllByCategory(categoryId, pageNo, pageNum)
	// }
	projects, err := StoreM.Project.FindWith(querySQL, pageNo, pageNum)
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
	pidi, err := strconv.Atoi(pids)
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

// func (p *Projects) Search(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
// 	query := req.URL.Query()
// 	tagId, _ := strconv.ParseInt(query.Get("tag_id"), 10, 64)
// 	// key := query.Get("key")

// 	projects, err := StoreM.Project.FindWithTag(tagId, 0, 10)
// 	if err != nil {
// 		return nil, DBErr
// 	}
// 	output, err := json.Marshal(WrapRespData(projects))
// 	if err != nil {
// 		log.Warnln("projects marshal projects,", err)
// 		return output, BadRespErr
// 	}
// 	return output, nil
// }
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
