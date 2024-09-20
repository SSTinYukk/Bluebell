package router

import (
	"bluebell/controller"
	"bluebell/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//创建
	r := gin.New()
	//中间件
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	//路由
	r.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
		defer c.Writer.CloseNotify()
		defer c.Request.Body.Close()
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
