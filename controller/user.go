package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bluebell/logic"
)

func SignUpHandler(c *gin.Context) {
	//1.验证参数

	//2.处理登录业务
	logic.SignUp()
	//3.返回响应
	c.JSON(http.StatusOK, "ok")
}
