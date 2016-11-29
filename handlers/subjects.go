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
