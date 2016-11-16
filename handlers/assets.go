package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"

	. "github.com/showntop/journey/stores"
)

type Assets struct {
	*application
}

func (a *Assets) List(req *http.Request) ([]byte, *HttpError) {
	assets, err := StoreM.Asset.FindAllWith("boot")
	if err != nil {
		log.Errorln("assets database error", err)
		return nil, DBErr
	}

	output, err := json.Marshal(WrapRespData(assets))
	if err != nil {
		log.Warnln("assets marshal assets,", err)
		return output, BadRespErr
	}
	return output, nil
}
