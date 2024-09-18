package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

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
	p.Password, err = hashPassword(p.Password)
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
