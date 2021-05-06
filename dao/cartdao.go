package dao

import (
	"fmt"
	"go_web/webapp/project_bookstore/model"
	"go_web/webapp/project_bookstore/utils"
)

//向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	//写sql
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	//获取购物车中所有的购物项
	cartItems := cart.CartItems
	//遍历得到每一个购物项
	for _, cartItem := range cartItems {
		//将购物项插入到数据库中
		AddCartItem(cartItem)
	}
	return err
}

//根据用户id去数据库查询对应的购物车
func GetCartByUserID(userid int) (*model.Cart, error) {
	//写sql语句
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id=?"
	//执行
	row := utils.Db.QueryRow(sqlStr, userid)
	//创建一个购物车
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	//获取当前购物车所有的购物项
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	//将所有的购物项设置到购物车中
	cart.CartItems = cartItems
	return cart, err
}

//更新购物车中的图书的总数量和总金额
func UpdateCart(cart *model.Cart) error {
	//写sql
	sqlStr := "update carts set total_count=?, total_amount=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)

	return err
}

//根据购物车的id删除购物车
func DeleteCartByCartID(cartID string) error {
	//删除购物车之前要删除所有的购物项
	err := DeleteCartItemsByCartID(cartID)
	if err != nil {
		fmt.Printf("DeleteCartItemsByCartID failed, err:%v\n", err)
		return err
	}
	//写sql语句
	sqlStr := "delete from carts where id=?"
	_, err = utils.Db.Exec(sqlStr, cartID)
	return err
}
