package entity

type ResponseObject struct {
	Code    int    `swaggertype:"string" example:"业务响应码（非HTTP码）"`
	Message string `example:"Some message about the code."`
}

func NewResponseObject(code int, message string) (r ResponseObject) {
	r.Code = code
	r.Message = message
	return
}
