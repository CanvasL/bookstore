package dao

import (
	"fmt"
	"go_web/webapp/project_bookstore/model"
	"go_web/webapp/project_bookstore/utils"
	"net/http"
)

//向数据库中添加Session
func AddSession(sess *model.Session) error {
	//写sql语句
	sqlStr := "insert into sessions values(?, ?, ?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	return err
}

//删除数据库中的Session
func DeleteSession(sessID string) error {
	//写sql语句
	sqlStr := "delete from sessions where session_id = ?"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, sessID)
	return err
}

//根据Session的Id值从数据库中查询Session
func GetSession(sessID string) (*model.Session, error) {
	//写sql语句
	sqlStr := "select session_id, username, user_id from sessions where session_id = ?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	//执行
	row := inStmt.QueryRow(sessID)
	//创建Session
	sess := &model.Session{}
	//扫描数据库中的字段值为Session的字段赋值
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, err
}

//判断用户是否已经登录,并返回session的UserName字段
func IsLogin(r *http.Request) (bool, *model.Session) {
	//根据Cookie的name获取Cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		fmt.Println("无用户登录")
	}
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//根据cookieValue去数据库中查询与之对应的Session
		session, err := GetSession(cookieValue)
		if err != nil {
			fmt.Printf("GetSession failed, err:%v\n", err)
		}
		if session.UserID > 0 {
			//已经登录
			return true, session
		}
	}
	return false, nil
}
