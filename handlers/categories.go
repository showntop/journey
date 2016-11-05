package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"

	. "github.com/showntop/journey/stores"
)

type Categories struct {
	*application
}

func (p *Categories) List(req *http.Request) ([]byte, *HttpError) {
	categories, err := StoreM.Category.FindAll()
	if err != nil {
		log.Errorln("categories database error", err)
		return nil, DBErr
	}

	output, err := json.Marshal(WrapRespData(categories))
	if err != nil {
		log.Warnln("categories marshal categories,", err)
		return output, BadRespErr
	}
	return output, nil
}
