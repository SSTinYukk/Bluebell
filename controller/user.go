package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"bluebell/logic"
	"bluebell/models"
)

func SignUpHandler(c *gin.Context) {
	//1.验证参数
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("signUp with invalid papram", zap.Error(err))
		//判断是不是validata error类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}
	if err := logic.SignUp(&p); err != nil {
		ResponseError(c, CodeUserExist)
		return
	}

	//3.返回响应
	ResponseSuccess(c, nil)

}

func LoginHandler(c *gin.Context) {
	//验证参数
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("login with invalid papram", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}
	//处理登录业务
	token, err := logic.Login(p)
	if err != nil {
		ResponseError(c, CodeUserNotExist)
		return
	}
	//返回处理结果
	ResponseSuccess(c, token)

}
