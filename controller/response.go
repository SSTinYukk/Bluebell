package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    StatusCode  ` json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResonseError(ctx *gin.Context, code StatusCode) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	})
}

func ResponseErrorWithMsg(ctx *gin.Context, code StatusCode, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    data,
	})
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	})
}
