package common

const (
	SUCCESS               = 200
	PARAMETER_BAD_REQUEST = 400
	UNAUTHORIZED          = 401
	FAIL                  = 500

	//用户(author) 相关状态码
	NAME_OR_PASSWORD_ERROR = 1001
	USER_NOT_FOUND         = 1002
	EMPTY_VISITOR_INFO     = 1003

	//文章Article 相关状态码
	ARTICLE_NOT_FOUND = 2001

	//label相关状态码
	EMPTY_LABEL_INFO = 3001
	LABEL_USED       = 3002

	//TOKEN 相关状态码
	TOKEN_WRONG_TYPE  = 9999
	TOKEN_WRONG_TOKEN = 9998
	TOKEN_NOT_VALID   = 9997
)

var Message = map[int]string{
	SUCCESS:               "OK",
	PARAMETER_BAD_REQUEST: "参数输入出现错误，请检查输入参数",
	UNAUTHORIZED:          "未授权操作，请登录或者确认自己的权限是否可以进行此操作",
	FAIL:                  "服务端出现错误，详情请查看日志",

	NAME_OR_PASSWORD_ERROR: "用户名或密码错误",
	USER_NOT_FOUND:         "未能找到该用户",
	EMPTY_VISITOR_INFO:     "用户名或密码为空",

	ARTICLE_NOT_FOUND: "找不到该文章",

	EMPTY_LABEL_INFO: "标签为空",
	LABEL_USED:       "标签名重复了",

	TOKEN_WRONG_TYPE:  "TOKEN格式错误（请不要伪造token）",
	TOKEN_WRONG_TOKEN: "TOKEN错误（请不要伪造token）",
	TOKEN_NOT_VALID:   "TOKEN无效（过期或未生效哦）",
}
