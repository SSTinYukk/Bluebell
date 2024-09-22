package controller

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/jwt"
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
	user, err := logic.Login(p)
	if err != nil {
		ResponseError(c, CodeUserNotExist)
		return
	}
	//返回处理结果
	ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", user.UserID),
		"username":      user.Username,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})

}
func RefreshTokenHandler(c *gin.Context) {
	//获取query参数
	rt := c.Query("refresh_token")
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "Token格式不对")
		c.Abort()
		return
	}
	//处理
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	if err != nil {
		zap.L().Error("jwt.RefreshToken failed", zap.Error(err))
		ResponseError(c, CodeInvalidToken)
		c.Abort()
		return
	}
	//返回

	ResponseSuccess(c, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
