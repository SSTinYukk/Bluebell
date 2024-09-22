package controller

import (
	"bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	//查询所有社区
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	//解析参数
	communityIdStr := c.Param("id")
	communityId, err := strconv.ParseUint(communityIdStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	community, err := logic.GetCommunityByID(communityId)
	if err != nil {
		zap.L().Error("logic.GetCommunityByID failed", zap.Error(err))
	}
	ResponseSuccess(c, community)
}
