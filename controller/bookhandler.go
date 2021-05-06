package controller

import (
	"fmt"
	"go_web/webapp/project_bookstore/dao"
	"go_web/webapp/project_bookstore/model"
	"net/http"
	"strconv"
	"text/template"
)

//IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中获取带分页的图书的函数
	page, err := dao.GetPageBooks(pageNo)
	if err != nil {
		fmt.Printf("GetPageBooks failed, err:%v\n", err)
	}
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)
}

//获取所有图书
//func GetBooks(w http.ResponseWriter, r *http.Request) {
//	//调用bookdao中获取所有图书的函数
//	books, err := dao.GetBooks()
//	if err != nil {
//		fmt.Printf("GetBooks failed, err:%v\n", err)
//	}
//	//解析模板文件
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	//执行
//	t.Execute(w, books)
//}

//获取带分页图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中获取带分页的图书的函数
	page, err := dao.GetPageBooks(pageNo)
	if err != nil {
		fmt.Printf("GetPageBooks failed, err:%v\n", err)
	}
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, page)
}

//获取带分页和价格范围的图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	//获取价格范围
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		//调用bookdao中获取带分页的图书的函数
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		//调用bookdao中获取带分页和价格范围的图书的函数
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		//将价格范围设置到page中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	//调用IsLogin函数判断是否已经登录
	flag, session := dao.IsLogin(r)

	if flag {
		//已经登录，设置page中的IsLogin字段和Username的字段值
		page.IsLogin = true
		page.Username = session.UserName
	}

	//解析模板文件
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)
}

////添加图书
//func AddBook(w http.ResponseWriter, r *http.Request) {
//	//获取图书信息
//	title := r.PostFormValue("title")
//	author := r.PostFormValue("author")
//	price, _:= strconv.ParseFloat(r.PostFormValue("price"), 64)
//	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
//	stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
//	//创建Book
//	book := &model.Book{
//		Title: title,
//		Author: author,
//		Price: price,
//		Sales: int(sales),
//		Stock: int(stock),
//		ImgPath: "static/img/default.jpg",
//	}
//	//调用bookdao中添加图书的函数
//	dao.AddBook(book)
//	//调用GetBooks函数再查询一次数据库
//	GetBooks(w, r)
//}

//删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//获取要删除的图书的id
	bookID := r.FormValue("bookId")
	//调用bookdao中删除图书的函数
	dao.DeleteBook(bookID)
	//调用GetBooks函数再查一次数据库
	GetPageBooks(w, r)
}

//去更新/添加图书的页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	//获取要更新的图书的id
	bookID := r.FormValue("bookId")
	//调用bookdao中获取图书的函数
	book, err := dao.GetBookByID(bookID)
	if err != nil {
		fmt.Printf("GetBookByID failed, err:%v\n", err)
	}
	if book.ID > 0 {
		//在更新图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, book)
	} else {
		//在添加图书
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, "")
	}

}

//更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	//获取图书信息
	bookID, _ := strconv.ParseInt(r.PostFormValue("bookId"), 10, 0)
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
	stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
	//创建Book
	book := &model.Book{
		ID:      int(bookID),
		Title:   title,
		Author:  author,
		Price:   price,
		Sales:   int(sales),
		Stock:   int(stock),
		ImgPath: "static/img/" + title + ".jpg",
	}
	if book.ID > 0 {
		//在更新图书
		//调用bookdao中更新图书的函数
		dao.UpdateBook(book)
	} else {
		//在添加图书
		//调用bookdao中添加图书的函数
		dao.AddBook(book)
	}
	//调用GetBooks函数再查询一次数据库
	GetPageBooks(w, r)
}
