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
	categoryId, _ := strconv.Atoi(queryValues.Get("category_id"))
	pageNo, _ := strconv.Atoi(queryValues.Get("page_no"))
	pageNum, _ := strconv.Atoi(queryValues.Get("page_num"))
	ptype := queryValues.Get("scope")
	_ = ptype

	var projects []*models.Project
	var err error
	if categoryId <= 0 {
		projects, err = StoreM.Project.FindAll(pageNo, pageNum)
	} else {
		projects, err = StoreM.Project.FindAllByCategory(categoryId, pageNo, pageNum)
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
