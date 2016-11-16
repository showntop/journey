package handlers

import ()

type application struct {
}

type HttpError struct {
	Code    int
	Message string
}

func (h *HttpError) Error() string {
	return h.Message
}

var (
	BadRequestErr       = &HttpError{400, "json format error"}
	BadRequestErr2      = &HttpError{401, "invalid query key/value pair"}
	IncorrectAccountErr = &HttpError{402, "用户名或者密码错误"}

	ServerErr  = &HttpError{500, "server error"}
	DBErr      = &HttpError{503, "db error"}
	BadRespErr = &HttpError{503, "response format error"}
)

func WrapRespData(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	result["data"] = data
	result["state_code"] = 200
	result["message"] = "成功"
	return result
}
