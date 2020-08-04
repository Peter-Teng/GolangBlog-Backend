package common

const (
	SUCCESS               = 200
	PARAMETER_BAD_REQUEST = 400
	FAIL                  = 500

	//Visitor 失败状态码
	NICKNAME_USED          = 1001
	NAME_OR_PASSWORD_ERROR = 1002
	VISITOR_DISABLED       = 1003
	VISITOR_NOT_FOUND      = 1004
	EMPTY_VISITOR_INFO     = 1005
)

var Message = map[int]string{
	SUCCESS:               "OK",
	PARAMETER_BAD_REQUEST: "参数输入出现错误，请检查输入参数",
	FAIL:                  "服务端出现错误，详情请查看日志",

	NICKNAME_USED:          "用户名已被使用",
	NAME_OR_PASSWORD_ERROR: "用户名或密码错误",
	VISITOR_DISABLED:       "用户已被禁用",
	VISITOR_NOT_FOUND:      "未能找到该用户",
	EMPTY_VISITOR_INFO:     "用户名或密码为空",
}
