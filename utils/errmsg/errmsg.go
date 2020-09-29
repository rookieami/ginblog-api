package errmsg

//错误处理模块
//状态码
const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	//code =1000...用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008

	//code=2000...分类模块的错误
	ERROR_CATEGORYNAME_USED = 2001
	ERROR_CATE_NOT_EXIST    = 3002

	//code=3000... 文章模块的错误
	ERROR_ART_NOT_EXIST = 3001
)

//状态码字典
var codeMsg = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN 不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式不正确",
	ERROR_USER_NO_RIGHT:    "该用户没权限",

	ERROR_CATEGORYNAME_USED: "分类名称已存在",
	ERROR_CATE_NOT_EXIST:    "该分类不存在",

	ERROR_ART_NOT_EXIST: "文章不存在",
}

func GetErrMsg(code int) string {
	msg, ok := codeMsg[code]
	if ok {
		return msg
	}
	return codeMsg[ERROR]
}
