package main

import (
	"Bases/5interface/discount"
	"fmt"
)

func main() {
	fmt.Println("11111111111111")
	order := &discount.Order{
		Id:    999,
		Total: 330.0,
	}
	PrintOrderDiscount(discount.FreeUser{Order: order})
	PrintOrderDiscount(discount.VipUser{Order: order})
}

type Discounter interface { // 定义接口
	ComputeOrder() //计算订单 没有返回值
}

func PrintOrderDiscount(d Discounter) { // 定义函数打印 订单金额  使用interface 计算不同的折扣  调用接口 Discounter
	d.ComputeOrder()
}
