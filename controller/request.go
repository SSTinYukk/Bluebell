package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrorUserNotlogin = errors.New("当前用户未登录")
)

const CtxUserIDKey = "userID"

func getCurrentUserID(c *gin.Context) (userID uint64, err error) {
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

func getPageInfo(c *gin.Context) (page, size int64, err error) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	if page, err = strconv.ParseInt(pageStr, 10, 64); err != nil {
		page = 1
	}
	if size, err = strconv.ParseInt(sizeStr, 10, 64); err != nil {
		size = 10
	}
	return
}
