package controller

import (
	"net/http"

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
		zap.L().Error("SignUp with invalid papram", zap.Error(err))
		//判断是不是validata error类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//2.处理登录业务
	logic.SignUp(&p)
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"user": p.Username,
	})

}
