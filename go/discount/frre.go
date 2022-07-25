package discount

import "fmt"

// 免费会员
type Order struct { //定义订单结构体
	Id    int
	Total float64
}

// 免费会员
type FreeUser struct { // 定义用户结构体
	Order *Order
}

func (u FreeUser) ComputeOrder() {
	fmt.Printf("订单编号：%d,订单金额：%2f\n", u.Order.Id, u.Order.Total*0.9)
}
