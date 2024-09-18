package mysql

import (
	"bluebell/models"
)

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
	sqlStr := "insert into user (user_id,username,password,email,gender) values(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Email, user.Gender)
	return
}
