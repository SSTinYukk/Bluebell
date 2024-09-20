package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrorUserNotlogin = errors.New("当前用户未登录")
)

const CtxUserIDKey = "userID"

func GetCurrentUserID(c *gin.Context) (userID uint64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotlogin
		return
	}
	userID, ok = uid.(uint64)
	if !ok {
		err = ErrorUserNotlogin
		return
	}
	return
}
