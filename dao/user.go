package dao

import (
	"suning/model"
)

// SearchUserByUsername 根据用户名查找用户
func SearchUserByUsername(username string) (u model.User, err error) {
	row := DB.QueryRow("select * from user where username=?", username)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.UserId, &u.Username, &u.Password)
	return
}

// InsertUser 在user表中添加数据
func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("insert into user (username,password) values (?,?)", u.Username, u.Password)
	return
}
