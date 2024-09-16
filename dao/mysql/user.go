package mysql

func CheckUserExist(userName string) bool {
	sqlStr := "select count(user_id) from user where username = ? "

	return 
}

func InsertUser() {
	//执行SQL语句入库

}
