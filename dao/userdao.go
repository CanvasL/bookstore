package dao

import (
	"go_web/webapp/project_bookstore/model"
	"go_web/webapp/project_bookstore/utils"
)

//验证用户名和密码
//根据用户名和密码在数据库中查询记录
func CheckUsernameAndPassword(username string, password string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id, username, password, email from users where username = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, err
}

//检查用户名是否已经存在
func CheckUsername(username string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id, username, password, email from users where username = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, err
}

//向数据库中插入用户信息
func SaveUser(username string, password string, email string) error {
	//写sql语句
	sqlStr := "insert into users(username, password, email) values(?, ?, ?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	return err
}
