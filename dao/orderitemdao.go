package dao

import (
	"go_web/webapp/project_bookstore/model"
	"go_web/webapp/project_bookstore/utils"
)

//向数据库中插入订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	//写sql语句
	sqlStr := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	return err
}

//根据订单号来获取所有的订单项
func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	//写sql
	sqlStr := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id=?"
	//执行
	rows, err := utils.Db.Query(sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderID)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
