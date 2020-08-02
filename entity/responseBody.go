package entity

type ResponseBody struct {
	Code    int
	Message string `example:"Some message about the code."`
}

func NewResponseBody(code int, message string) (r ResponseBody) {
	r.Code = code
	r.Message = message
	return
}
