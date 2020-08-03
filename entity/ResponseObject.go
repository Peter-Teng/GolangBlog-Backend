package entity

type ResponseObject struct {
	Code    int
	Message string `example:"Some message about the code."`
}

func NewResponseObject(code int, message string) (r ResponseObject) {
	r.Code = code
	r.Message = message
	return
}
