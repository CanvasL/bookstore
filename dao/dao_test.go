package dao

import (
	"fmt"
	"go_web/webapp/project_bookstore/model"
	"testing"
	"time"
)

func testUser(t *testing.T) {
	fmt.Println("测试userdao中的函数...")
	//t.Run("验证用户名或密码：", testLogin)
	//t.Run("验证用户名：", testRegist)
	//t.Run("保存用户：", testSave)
}

func testBook(t *testing.T) {
	fmt.Println("测试bookdao中的相关函数")
	//t.Run("测试获取所有图书", testGetBooks)
	//t.Run("测试添加图书", testAddBook)
	//t.Run("测试删除图书", testDeleteBook)
	//t.Run("测试通过ID获取图书", testGetBookByID)
	//t.Run("测试更新图书", testUpdateBook)
	//t.Run("测试获取带分页的图书", testGetPageBooks)
	//t.Run("测试获取带分页和价格范围的图书", testGetPageBooksByPrice)
}

func testSession(t *testing.T) {
	fmt.Println("测试sessiondao中的相关函数")
	//t.Run("测试添加Session", testAddSession)
	//t.Run("测试删除Session", testDeleteSession)
	//t.Run("测试从数据库中查询到Session", testGetSession)
}

func testCart(t *testing.T) {
	fmt.Println("测试cartdao中的相关函数")
	//t.Run("测试添加Cart", testAddCart)
	//t.Run("测试根据userid查询表项", testGetCartByUserID)
	//t.Run("更新Cart表项", testUpdateCart)
	t.Run("测试根据cartID删除cart", testDeleteCartItemsByCartID)
}

func testCartItem(t *testing.T) {
	fmt.Println("测试cartitemdao中的相关函数")
	//t.Run("测试根据bookID查询表项", TestGetCartItemByBookIDAndCartID)
	//t.Run("测试根据cartID查询表项", testGetCartItemsByCartID)
	//t.Run("测试根据bookID和cartID更新图书数量", testUpdateBookCount)
	t.Run("根据ID来删除购物项", testDeleteCartItemByID)
}

func TestOrder(t *testing.T) {
	fmt.Println("测试orderdao中的相关函数")
	//t.Run("测试添加订单", testAddOrder)
	//t.Run("测试拿到所有订单", testGetOrders)
	//t.Run("测试获取我的订单", testGetMyOrders)
	t.Run("测试更新订单的状态(发货和收货)", testUpdateOrderState)
}

func testOrderItem(t *testing.T) {
	fmt.Println("测试orderitemdao中的相关函数")
	//t.Run("根据ID来删除购物项", testDeleteCartItemByID)
	t.Run("根据订单id来获取所有订单项", testGetOrderItemsByOrderID)
}

//userdao============================================================================

func testLogin(t *testing.T) {
	user, err := CheckUsernameAndPassword("admin", "123456")
	if err != nil {
		fmt.Println("testLogin err:", err)
	} else {
		fmt.Println("获取的用户信息是：", user)
	}
}

func testRegist(t *testing.T) {
	user, err := CheckUsername("admin")
	if err != nil {
		fmt.Println("testRegister err:", err)
	} else {
		fmt.Println("用户名已存在：", user)
	}
}

func testSave(t *testing.T) {
	SaveUser("admin0", "1234567", "100001@qq.com")
}

//bookdao============================================================================
func testGetBooks(t *testing.T) {
	books, err := GetBooks()
	if err != nil {
		fmt.Printf("GetBooks failed, err:%v\n", err)
	}
	//遍历每一本图书
	for k, v := range books {
		fmt.Printf("第%v本图书的信息是:%v\n", k+1, v)
	}
}

func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "static/img/default.jpg",
	}
	//调用添加图书的函数
	AddBook(book)
}

func testDeleteBook(t *testing.T) {
	//调用删除图书的函数
	DeleteBook("34")
}

func testGetBookByID(t *testing.T) {
	//调用后获取图书的函数
	book, err := GetBookByID("38")
	if err != nil {
		fmt.Printf("GetBookByID failed, err:%v\n", err)
	}
	fmt.Println("获取的图书信息是：", book)
}

func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      31,
		Title:   "三国志",
		Author:  "陈寿",
		Price:   49.80,
		Sales:   5,
		Stock:   168,
		ImgPath: "static/img/default.jpg",
	}
	//调用更新图书的函数
	err := UpdateBook(book)
	if err != nil {
		fmt.Printf("UpdateBook failed, err:%v\n", err)
	}
}

func testGetPageBooks(t *testing.T) {
	page, err := GetPageBooks("9")
	if err != nil {
		fmt.Printf("GetPageBooks failed, err:%v\n", err)
	}
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("当记录数是：", page.TotalRecord)
	fmt.Println("当前页的图书有：")
	for _, v := range page.Books {
		fmt.Println("图书信息是：", v)
	}
}

func testGetPageBooksByPrice(t *testing.T) {
	page, err := GetPageBooksByPrice("1", "20", "30")
	if err != nil {
		fmt.Printf("GetPageBooks failed, err:%v\n", err)
	}
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("当记录数是：", page.TotalRecord)
	fmt.Println("当前页的图书有：")
	for _, v := range page.Books {
		fmt.Println("图书信息是：", v)
	}
}

//sessiondao============================================================================

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "123456677",
		UserName:  "马化腾",
		UserID:    2,
	}
	AddSession(sess)
}

func testDeleteSession(t *testing.T) {
	DeleteSession("123456677")
}

func testGetSession(t *testing.T) {
	sess, _ := GetSession("1fdc5615-3a89-40bc-60e8-7245a1744a15")
	fmt.Println("Session的信息是：", sess)
}

//cartdao============================================================================

func testAddCart(t *testing.T) {
	//设置要买的两本书Book
	book1 := &model.Book{
		ID:    1,
		Price: 27.20,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}
	//创建两个购物项CartItem
	cartItem1 := &model.CartItem{
		Book:   book1,
		Count:  10,
		CartID: "6668888",
	}
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "6668888",
	}
	//创建购物项切片
	var cartItems []*model.CartItem
	cartItems = append(cartItems, cartItem1)
	cartItems = append(cartItems, cartItem2)
	//创建购物车Cart
	cart := &model.Cart{
		CartID:    "6668888",
		CartItems: cartItems,
		UserID:    7,
	}
	//将购物车插入到数据库中
	AddCart(cart)
}

func testGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(7)
	fmt.Println("id为7的用户的购物车信息：", cart)
}

func testUpdateCart(t *testing.T) {
	book1 := &model.Book{
		ID:    1,
		Price: 27.20,
	}
	cartItem1 := &model.CartItem{
		Book:   book1,
		Count:  1,
		CartID: "10f96f83-468d-4a41-78d4-7d38d6d467d2",
	}
	//创建购物项切片
	var cartItems []*model.CartItem
	cartItems = append(cartItems, cartItem1)
	//创建购物车Cart
	cart := &model.Cart{
		CartID:    "10f96f83-468d-4a41-78d4-7d38d6d467d2",
		CartItems: cartItems,
		UserID:    1,
	}
	//更新cart
	UpdateCart(cart)
}

func testDeleteCartItemsByCartID(t *testing.T) {
	DeleteCartByCartID("84c47ba4-914d-42e2-666b-51fa835dfaba")
}

//cartitemdao============================================================================

func testGetCartItemByBookIDAndCartID(t *testing.T) {
	cartItem, _ := GetCartItemByBookIDAndCartID("1", "6668888")
	fmt.Println("图书ID=1的图书信息是：", cartItem)
}

func testGetCartItemsByCartID(t *testing.T) {
	cartItems, _ := GetCartItemsByCartID("6668888")
	for k, v := range cartItems {
		fmt.Printf("第%v个购物项是：%v\n", k+1, v)
	}
}

func testUpdateBookCount(t *testing.T) {
	//已验证
}

func testDeleteCartItemByID(t *testing.T) {
	err := DeleteCartItemByID("34")
	if err != nil {
		fmt.Printf("DeleteCartItemByID failed, err:%v\n", err)
	}
}

//orderdao============================================================================

func testAddOrder(t *testing.T) {
	//生成订单号
	orderID := "1231231321321321"
	//创建订单
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  2,
		TotalAmount: 400.00,
		State:       0, //未发货
		UserID:      8,
	}
	//创建订单项1
	orderItem1 := &model.OrderItem{
		Count:   1,
		Amount:  300.00,
		Title:   "300块钱能做什么？",
		Author:  "梁帆",
		Price:   300.00,
		ImgPath: "/static/img/default.jpg",
		OrderID: orderID,
	}
	//创建订单项2
	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100.00,
		Title:   "100块钱能做什么？",
		Author:  "梁帆",
		Price:   100.00,
		ImgPath: "/static/img/default.jpg",
		OrderID: orderID,
	}
	//保存订单
	err := AddOrder(order)
	if err != nil {
		fmt.Printf("AddOrder failed, err:%v\n", err)
	}
	//保存订单项
	err = AddOrderItem(orderItem1)
	if err != nil {
		fmt.Printf("AddOrderItem1 failed, err:%v\n", err)
	}
	AddOrderItem(orderItem2)
	if err != nil {
		fmt.Printf("AddOrderItem2 failed, err:%v\n", err)
	}
}

func testGetOrders(t *testing.T) {
	orders, _ := GetOrders()
	for _, v := range orders {
		fmt.Printf("订单信息是：%v\n", v)
	}
}

func testGetMyOrders(t *testing.T) {
	orders, _ := GetMyOrders(8)
	for _, v := range orders {
		fmt.Printf("我的订单有：%v\n", v)
	}
}

func testUpdateOrderState(t *testing.T) {
	err := UpdateOrderState("0fd072c1-0cc2-488a-474c-735f76562ab1", 1)
	if err != nil {
		fmt.Printf("UpdateOrderState failed, err:%v\n", err)
	}
}

//orderitemdao============================================================================
func testGetOrderItemsByOrderID(t *testing.T) {
	orderItems, _ := GetOrderItemsByOrderID("fd6a8f85-8aaf-40d1-5a89-ab07dfac5fd5")
	for _, v := range orderItems {
		fmt.Printf("订单项的信息是：%v\n", v)
	}
}
