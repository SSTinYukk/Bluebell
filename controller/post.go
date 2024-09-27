package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	//1.获取参数
	if err := c.ShouldBind(p); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID(c) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	p.AuthorId = userID

	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseUint(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		return
	}
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("GetPostByID(pid) failed", zap.Error(err))
	}
	ResponseSuccess(c, data)
}
func GetPostListHandler(c *gin.Context) {
	page, size, err := getPageInfo(c)
	if err != nil {
		zap.L().Error("getPageInfo(c) failed", zap.Error(err))
		return
	}
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("GetPostList() failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, data)
}
