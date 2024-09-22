package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户是否存在
	exist, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return
	}
	if exist {
		err = errors.New("user already exist")
		return
	}
	//生成UID
	userID, err := snowflake.GenID()
	if err != nil {
		return
	}
	user := &models.User{
		UserID:   userID,
		Password: p.Password,
		Username: p.Username,
		Email:    p.Email,
		Gender:   p.Gender,
	}
	if err = mysql.InsertUser(user); err != nil {
		return
	}
	//返回
	return nil
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err = mysql.VerifyPassword(user); err != nil {
		return
	}
	user.AccessToken, user.RefreshToken, err = jwt.GenToken(user.UserID, p.Username)
	return
}
