package mysql

import (
	"bluebell/models"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckUserExist(username string) (bool, error) {
	sqlStr := "select count(user_id) from user where username = ? "
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return false, nil
}

func InsertUser(user *models.User) (err error) {
	//执行SQL语句入库
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return
	}
	sqlStr := "insert into user (user_id,username,password,email,gender) values(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Email, user.Gender)
	return
}

func VerifyPassword(user *models.User) (err error) {
	pwd := user.Password
	sqlStr := "select user_id,username,password from user where username = ?"
	if err = db.Get(user, sqlStr, user.Username); err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))

	return
}
