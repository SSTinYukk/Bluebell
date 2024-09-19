package controller

type StatusCode int64

const (
	CodeSuccess           StatusCode = 1000
	CodeInvalidParams     StatusCode = 1001
	CodeUserExist         StatusCode = 1002
	CodeUserNotExist      StatusCode = 1003
	CodeInvalidPassword   StatusCode = 1004
	CodeServerBusy        StatusCode = 1005
	CodeInvalidToken      StatusCode = 1006
	CodeInvalidAuthFormat StatusCode = 1007
	CodeNotLogin          StatusCode = 1008
	ErrVoteRepeated       StatusCode = 1009
	ErrorVoteTimeExpire   StatusCode = 1010
)

var msgFlags = map[StatusCode]string{
	CodeSuccess:           "成功",
	CodeInvalidParams:     "参数错误",
	CodeUserExist:         "用户已存在",
	CodeUserNotExist:      "用户不存在",
	CodeInvalidPassword:   "用户名或密码错误",
	CodeServerBusy:        "服务繁忙",
	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",
	ErrVoteRepeated:       "请勿重复投票",
	ErrorVoteTimeExpire:   "投票时间已过",
}

func (c StatusCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
