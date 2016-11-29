package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/journey/handlers"
)

func WrapErrorResp(err *handlers.HttpError) []byte {
	output := []byte(`{
		"message": "response json error",
		"state_code": 503
		}`)
	output, _ = json.Marshal(map[string]interface{}{
		"state_code": err.Code,
		"message":    err.Message,
	})

	return output
}

func Instrument() *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/subjects", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		SubjectsC := new(handlers.Subjects)
		results, err := SubjectsC.List(req, ps)
		if err != nil {
			// http.Error(rw, err.Error(), err.Code)
			rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	router.GET("/api/v1/components", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		ComponentsC := new(handlers.Components)
		results, err := ComponentsC.List(req)
		if err != nil {
			// http.Error(rw, err.Error(), err.Code)
			rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	router.GET("/api/v1/assets", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		categoriesC := new(handlers.Assets)
		results, err := categoriesC.List(req)
		if err != nil {
			// http.Error(rw, err.Error(), err.Code)
			rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	router.GET("/api/v1/categories", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		assetsC := new(handlers.Categories)
		results, err := assetsC.List(req)
		if err != nil {
			// http.Error(rw, err.Error(), err.Code)
			rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	router.GET("/api/v1/apps", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		projectsC := new(handlers.Projects)
		results, err := projectsC.List(req)
		if err != nil {
			// http.Error(rw, err.Error(), err.Code)
			rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	router.GET("/api/v1/search", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		projectsC := new(handlers.Projects)
		results, err := projectsC.Search(req, ps)
		if err != nil {
			// http.Error(rw, err.Error(), err.Code)
			rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	router.GET("/api/v1/apps/:id", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		projectsC := new(handlers.Projects)
		results, err := projectsC.Show(req, ps)
		if err != nil {
			// http.Error(rw, err.Error(), err.Code)
			rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	router.POST("/api/v1/users", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		usersC := new(handlers.Users)
		results, err := usersC.Create(req)
		if err != nil {
			// rw.Write([]byte(err.Error()))
			rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	router.POST("/api/v1/sessions", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		sessionsC := new(handlers.Sessions)
		results, err := sessionsC.Create(req)
		if err != nil {
			// http.Error(rw, err.Error(), err.Code)
			rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	return router
}
