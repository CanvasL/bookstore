package dao

import (
	"go_web/webapp/project_bookstore/model"
	"go_web/webapp/project_bookstore/utils"
)

//插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	//写sql
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)

	return err
}

//根据BookID去查询购物项
func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	//写sql
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id=? and cart_id=?"
	//执行
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	//创建cartItem
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	//根据图书的id获取对应的图书信息
	book, _ := GetBookByID(bookID)
	//将book设置到购物项中
	cartItem.Book = book
	return cartItem, nil
}

//根据购物项相关信息更新购物项中图书的数量
func UpdateBookCount(cartItem *model.CartItem) error {
	//写sql
	sqlStr := "update cart_items set count=?, amount=? where book_id=? and cart_id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	return err
}

//根据CartID去查询所有的购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	//写sql
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id=?"
	//执行
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		//设置一个变量接收bookID
		var bookID string
		//创建cartItem
		cartItem := &model.CartItem{}
		err = rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err != nil {
			return nil, err
		}
		//根据bookID获取图书信息
		book, _ := GetBookByID(bookID)
		//将book设置到购物项中
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, err
}

//根据购物车的id来删除所有的购物项
func DeleteCartItemsByCartID(cartID string) error {
	//写sql
	sqlStr := "delete from cart_items where cart_id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartID)
	return err
}

//根据购物项id来删除购物项
func DeleteCartItemByID(cartItemID string) error {
	//写sql语句
	sqlStr := "delete from cart_items where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItemID)
	return err
}
