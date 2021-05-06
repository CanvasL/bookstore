package controller

import (
	"fmt"
	"go_web/webapp/project_bookstore/dao"
	"go_web/webapp/project_bookstore/model"
	"go_web/webapp/project_bookstore/utils"
	"net/http"
	"text/template"
)

//Login处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登录
	flag, _ := dao.IsLogin(r)
	if flag {
		//已经登陆过了
		//去首页
		GetPageBooksByPrice(w, r)
	}
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//调用userdao中验证用户名和密码的方法
	user, err := dao.CheckUsernameAndPassword(username, password)
	if err != nil {
		fmt.Printf("CheckUsernameAndPassword failed, err:%v\n", err)
	}
	if user.ID > 0 {
		//用户名和密码正确
		//生成UUID作为Session的id
		uuid := utils.CreateUUID()
		//创建一个session
		sess := &model.Session{
			SessionID: uuid,
			UserName:  user.Username,
			UserID:    user.ID,
		}
		//将Session保存到数据库中
		dao.AddSession(sess)
		//创建Cookie，让它与Session关联
		cookie := http.Cookie{
			Name:     "user",
			Value:    uuid,
			HttpOnly: true,
		}
		//将Cookie发送给浏览器
		http.SetCookie(w, &cookie)
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, user)
	} else {
		//用户名或密码不正确
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	}
}

//处理用户注销的函数
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, err := r.Cookie("user")
	if err != nil {
		fmt.Printf("获取cookie失效，err:%v\n", cookie)
	}
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中与之对应的Session
		dao.DeleteSession(cookieValue)
		//设置Cookie失效
		cookie.MaxAge = -1
		//将修改之后的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	//去首页
	GetPageBooksByPrice(w, r)
}

//Regist处理用户注册的数据
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取用户名、密码和邮箱
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//调用userdao中的CheckUsername方法
	user, err := dao.CheckUsername(username)
	if err != nil {
		fmt.Printf("CheckUsername failed, err:%v\n", err)
	}
	if user.ID > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		//用户名可用
		err = dao.SaveUser(username, password, email)
		if err != nil {
			fmt.Printf("SaveUser failed, err:%v\n", err)
		}
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

//通过发送Ajax请求验证用户名
func CheckUsername(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的用户名
	username := r.PostFormValue("username")
	user, err := dao.CheckUsername(username)
	if err != nil {
		fmt.Printf("CheckUsername failed, err:%v\n", err)
	}
	if user.ID > 0 {
		//用户名已存在
		w.Write([]byte("用户名已存在！"))
	} else {
		//用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}
