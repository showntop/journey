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

type Subjects struct {
	*application
}

func (p *Subjects) List(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	queryValues := req.URL.Query()
	pageNo, _ := strconv.Atoi(queryValues.Get("page_no"))
	pageNum, _ := strconv.Atoi(queryValues.Get("page_num"))

	var subjects []*models.Subject
	var err error
	subjects, err = StoreM.Subject.FindAll(pageNo, pageNum)
	if err != nil {
		log.Errorln("subjects database error", err)
		return nil, DBErr
	}
	output, err := json.Marshal(WrapRespData(subjects))
	if err != nil {
		log.Warnln("subjects marshal subjects,", err)
		return output, BadRespErr
	}
	return output, nil
}

func (p *Subjects) Show(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {

	key := "植物"
	sid, _ := strconv.ParseInt(ps.ByName("id"), 10, 64)
	subject, err := StoreM.Subject.Find(sid)
	if err != nil {
		return nil, DBErr
	}

	key = subject.Params

	projects, err := StoreM.Project.FindWithKey(key, 0, 6)
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
