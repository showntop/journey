package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"

	. "github.com/showntop/journey/stores"
)

type Sessions struct {
}

func (s *Sessions) Create(req *http.Request) ([]byte, *HttpError) {
	lgoinInfo := &struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&lgoinInfo)
	if err != nil {
		log.Warnln("session parse request,", err)
		return nil, BadRequestErr
	}

	user, err := StoreM.User.FindBy(lgoinInfo.Username)
	if err != nil {
		log.Errorln("session query database,", err)
		return nil, DBErr
	}
	if user == nil {
		return nil, IncorrectAccountErr
	}
	user.Password = lgoinInfo.Password
	if err := user.Authenticate(); err != nil {
		log.Warnln("session auth user,", err)
		return nil, IncorrectAccountErr
	}

	//respose do
	output, err := json.Marshal(WrapRespData(user))
	if err != nil {
		log.Warnln("session marshal user,", err)
		return output, BadRespErr
	}
	return output, nil
}
