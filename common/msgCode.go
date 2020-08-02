package common

const (
	SUCCESS = 200
	FAIL    = 500

	//Visitor 失败状态码
	NICKNAME_USED          = 1001
	NAME_OR_PASSWORD_ERROR = 1002
	VISITOR_DISABLED       = 1003
	VISITOR_NOT_FOUND      = 1004
)

var Message = map[int]string{
	SUCCESS: "OK",
	FAIL:    "服务端出现错误，详情请查看日志",

	NICKNAME_USED:          "用户名已被使用",
	NAME_OR_PASSWORD_ERROR: "用户名或密码错误",
	VISITOR_DISABLED:       "用户已被禁用",
	VISITOR_NOT_FOUND:      "未能找到该用户",
}