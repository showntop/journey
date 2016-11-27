package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/showntop/journey/models"
	. "github.com/showntop/journey/stores"
)

type Users struct {
}

func (u *Users) Create(req *http.Request) ([]byte, *HttpError) {
	//request do
	signupInfo := &struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Mobile   string `json:"mobile"`
	}{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&signupInfo)
	if err != nil {
		return nil, BadRequestErr
	}
	//only allowed name field
	//did has the better way or separate into reqmodel  sqlmodel  repmodel
	user := &models.User{
		Username: signupInfo.Username,
		Password: signupInfo.Password,
		Mobile:   signupInfo.Mobile,
		Email:    signupInfo.Email,
	}
	if err := user.EncryptPassword(); err != nil {
		return nil, ServerErr
	}
	if verr := user.Validate(); verr != nil {
		return nil, BadRequestErr
	}

	//save
	err = StoreM.User.Create(user)
	if err != nil {
		log.Error("server error:", err)
		return nil, DBErr
	}

	//respose do
	output, err := json.Marshal(WrapRespData(user))
	if err != nil {
		return output, BadRespErr
	}
	return output, nil
}
