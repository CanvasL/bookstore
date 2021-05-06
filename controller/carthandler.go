package controller

import (
	json2 "encoding/json"
	"fmt"
	"go_web/webapp/project_bookstore/dao"
	"go_web/webapp/project_bookstore/model"
	"go_web/webapp/project_bookstore/utils"
	"net/http"
	"strconv"
	"text/template"
)

//添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, session := dao.IsLogin(r)
	if flag {
		//已经登录
		//获取要添加的图书的id
		bookID := r.FormValue("bookId")
		//根据图书的id来获取图书信息
		book, err := dao.GetBookByID(bookID)
		if err != nil {
			fmt.Printf("GetBookByID failed, err:%v\n", err)
		}
		//获取用户的id
		userID := session.UserID
		//判断数据库中是否有当前用户的购物车
		cart, err := dao.GetCartByUserID(userID)
		if err != nil {
			fmt.Printf("GetCartByUserID failed, err:%v\n", err)
		}
		if cart != nil {
			//当前用户已经有购物车，此时需要判断购物车中是否有当前这本图书
			cartItem, err := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			if err != nil {
				//fmt.Printf("GetCartItemByBookIDAndCartID failed, err:%v\n", err)
				fmt.Println("当前购物车中还没有该图书对应的购物项")
			}
			if cartItem != nil {
				//购物车中购物项已有该图书，只需要将其数量+1即可
				fmt.Println("用户已经有该图书")
				//1.获取购物车切片中的所有的购物项
				cts := cart.CartItems
				//2.遍历得到每一个购物项
				for _, v := range cts {
					//3.找到当前的购物项
					if cartItem.Book.ID == v.Book.ID {
						//将当前购物项中图书的数量+1
						v.Count++
						//将数据库中cart_items的数量更新
						dao.UpdateBookCount(v)
						//更新carts表项的操作在后
					}
				}
			} else {
				//购物车的购物项中还没有该图书，此时需要创建一个购物项并添加到数据库中
				//创建购物车中的购物项
				cartItem = &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}

				//将购物项添加到当前cart的切片中
				cart.CartItems = append(cart.CartItems, cartItem)
				//将新建的购物项添加到数据库中
				err = dao.AddCartItem(cartItem)
				if err != nil {
					fmt.Printf("AddCart failed, err:%v\n", err)
				}
			}
			//不管之前购物车中是否有当前图书所对应的购物项，都需要更新购物车中的图书的总数量和总金额
			dao.UpdateCart(cart)
		} else {
			//证明当前用户还没有购物车，需要创建一个购物车并添加到数据库中
			//1.创建购物车
			cartID := utils.CreateUUID()
			cart := &model.Cart{
				CartID: cartID,
				UserID: userID,
			}
			//2.创建购物车中的购物项
			//声明一个切片
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cartID,
			}
			//3.将购物项添加到切片中
			cartItems = append(cartItems, cartItem)
			//4.将切片添加到购物车中
			cart.CartItems = cartItems
			//5.将购物车cart保存到数据库中
			err = dao.AddCart(cart)
			if err != nil {
				fmt.Printf("AddCart failed, err:%v\n", err)
			}
		}
		w.Write([]byte("您刚刚将《" + book.Title + "》添加到了购物车!"))
	} else {
		//没有登录
		w.Write([]byte("请求登录！"))
	}
}

//获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	//获取用户的id
	userID := session.UserID
	//根据用户的id从数据库中获取对应的购物车
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		//将购物车设置到session中
		session.Cart = cart
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, session)
	} else {
		//该用户还没有购物车
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, session)
	}
}

//清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物车的id
	cartID := r.FormValue("cartId")
	//清空购物车
	err := dao.DeleteCartByCartID(cartID)
	if err != nil {
		fmt.Printf("DeleteCartByCartID failed, err:%v\n", err)
	}
	//调用GetCartInfo再次查询购物车信息
	GetCartInfo(w, r)
}

//删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物项id
	cartItemID := r.FormValue("cartItemId")
	//将购物项的id转换成int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取session
	_, session := dao.IsLogin(r)
	//获取用户的id
	userID := session.UserID
	//获取用户的购物车
	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		fmt.Printf("GetCartByUserID failed, err：%v\n", err)
	}
	//获取购物车的所有购物项
	cartItems := cart.CartItems
	//遍历得到每一个购物项
	for k, v := range cartItems {
		//寻找要删除的购物项
		if v.CartItemID == iCartItemID {
			//这个就是我们要删除的购物项
			//将当前购物项从切片中移除
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			//将删除购物项之后的切片再次赋给购物车中的切片
			cart.CartItems = cartItems
			//将当前购物项从数据库中移除
			dao.DeleteCartItemByID(cartItemID)
		}
	}
	//更新购物车中的图书的总数量和总金额
	dao.UpdateCart(cart)
	//再次调用获取购物车信息的函数
	GetCartInfo(w, r)
}

//更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要更新的购物项id,将购物项的id转换成int64
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取用户输入的图书数量,将数量转换成int64
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	//获取session
	_, session := dao.IsLogin(r)
	//获取用户的id
	userID := session.UserID
	//获取用户的购物车
	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		fmt.Printf("GetCartByUserID failed, err：%v\n", err)
	}
	//获取购物车的所有购物项
	cartItems := cart.CartItems
	//遍历得到每一个购物项
	for _, v := range cartItems {
		//寻找要更新的购物项
		if v.CartItemID == iCartItemID {
			//这个就是我们要更新的购物项
			v.Count = iBookCount
			//更新数据库中该购物项的图书数量和金额小计
			err = dao.UpdateBookCount(v)
			if err != nil {
				fmt.Printf("UpdateBookCount failed, err:%v\n", err)
			}
		}
	}
	//更新购物车中的图书的总数量和总金额
	dao.UpdateCart(cart)
	//调用获取购物项信息的函数再次查询购物车信息
	cart, _ = dao.GetCartByUserID(userID)
	//再次调用获取购物车信息的函数
	//GetCartInfo(w, r)
	//获取购物车中图书的总数量
	totalCount := cart.TotalCount
	//获取购物车中图书的总金额
	totalAmount := cart.TotalAmount
	var amount float64
	//获取购物车中更新的购物项中的金额小计
	cIs := cart.CartItems
	for _, v := range cIs {
		if iCartItemID == v.CartItemID {
			//这就是我们要的购物项，则获取当前购物项的金额小计
			amount = v.Amount
		}
	}
	//创建Data结构
	data := model.Data{
		TotalCount:  totalCount,
		TotalAmount: totalAmount,
		Amount:      amount,
	}
	//将data转换为json字符串
	json, _ := json2.Marshal(data)
	//响应到浏览器
	w.Write(json)
}
