package controller

import (
	"bluebell/logic"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UploadResourceHandler 处理文件上传请求
func UploadResourceHandler(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		zap.L().Error("获取上传文件失败", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	// 读取文件内容
	fileData, err := file.Open()
	if err != nil {
		zap.L().Error("打开上传文件失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer fileData.Close()

	data, err := io.ReadAll(fileData)
	if err != nil {
		zap.L().Error("读取上传文件内容失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 获取当前用户 ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("获取当前用户 ID 失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 处理文件上传逻辑
	resource, err := logic.UploadResource(data, file.Filename, userID)
	if err != nil {
		zap.L().Error("文件上传逻辑处理失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回成功响应
	ResponseSuccess(c, resource)
}

// GetResourceByIDHandler 根据资源 ID 获取文件资源信息
func GetResourceByIDHandler(c *gin.Context) {
	// 获取资源 ID
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("解析资源 ID 失败", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	// 获取文件资源信息
	resource, err := logic.GetResourceByID(id)
	if err != nil {
		zap.L().Error("获取文件资源信息失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回成功响应
	ResponseSuccess(c, resource)
}
