package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake.go"
)

func SignUp(p *models.ParamSignUp) {
	//判断用户是否存在
	mysql.CheckUserExist(p.Username)

	//生成UID
	snowflake.GenID()
	//返回
}
