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

type Posts struct {
	application
}

func (p *Posts) List(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	queryValues := req.URL.Query()
	projectId, _ := strconv.Atoi(ps.ByName("id"))
	pageNo, _ := strconv.Atoi(queryValues.Get("page_no"))
	pageNum, _ := strconv.Atoi(queryValues.Get("page_num"))

	var posts []*models.Post
	var err error
	posts, err = StoreM.Post.FindByProject(projectId, pageNo, pageNum)
	if err != nil {
		log.Errorln("posts database error", err)
		return nil, DBErr
	}
	output, err := json.Marshal(WrapRespData(posts))
	if err != nil {
		log.Warnln("posts marshal posts,", err)
		return output, BadRespErr
	}
	return output, nil
}

func (u *Posts) Create(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	user, err := u.AuthUser(req)
	if err != nil {
		return nil, InvalidTokenErr
	}
	//request do
	var post models.Post
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&post)
	if err != nil {
		return nil, BadRequestErr
	}
	//only allowed name field
	//did has the better way or separate into reqmodel  sqlmodel  repmodel
	//post.valid
	//save
	post.UserId = user.Id
	post.ProjectId, _ = strconv.ParseInt(ps.ByName("id"), 10, 64)
	err = StoreM.Post.Create(&post)

	if err != nil {
		log.Errorln("handle posts create, message:", err)
		return nil, DBErr
	}

	//respose do
	output, err := json.Marshal(WrapRespData(post))
	if err != nil {
		return output, BadRespErr
	}
	return output, nil
}

func (p *Posts) CreateLike(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	//request do
	var postLike models.PostLike
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&postLike)
	if err != nil {
		return nil, BadRequestErr
	}
	//only allowed name field
	//did has the better way or separate into reqmodel  sqlmodel  repmodel
	//post.valid
	//save
	err = StoreM.Post.CreateLike(&postLike)
	if err != nil {
		return nil, DBErr
	}

	//respose do
	output, err := json.Marshal(WrapRespData(postLike))
	if err != nil {
		return output, BadRespErr
	}
	return output, nil
}

func (p *Posts) CreateComment(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	//request do
	var postComment models.PostComment
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&postComment)
	if err != nil {
		return nil, BadRequestErr
	}
	//only allowed name field
	//did has the better way or separate into reqmodel  sqlmodel  repmodel
	//post.valid
	//save
	err = StoreM.Post.CreateComment(&postComment)
	if err != nil {
		return nil, DBErr
	}

	//respose do
	output, err := json.Marshal(WrapRespData(postComment))
	if err != nil {
		return output, BadRespErr
	}
	return output, nil
}
