package main

import "fmt"

func main() {
	// 条件控制语句  (分支控制)
	//  if    else if     else
	//  switch
	// fallthrough

	//需求 计算订单优惠价
	// 1 订单金额 小于等于100 打9 折
	// 2 订单金额 大于100等于300  打8 折
	// 3 订单金额 大于300 打5折

	A := 20.00
	Y := 0.0
	Z := 0.0
	if A <= 100 { //   go语言中 if 后面的括号 不写
		// 1 订单金额 小于等于100 打9 折
		Y = A * 0.9

	} else if A > 100 && A <= 300 {
		// 2 订单金额 大于100等于300  打8 折
		Y = A * 0.8
	} else {
		// 2 订单金额 大于100等于300  打8 折
		Y = A * 0.5

	}
	fmt.Printf("折扣后的价格为：%.2f\n", Y) // %.2f 保留浮点数两位

	switch {

	case A <= 100:
		Z = A * 0.9
	case A > 100 && A <= 300:
		Z = A * 0.8
	default:
		Z = A * 0.5

	}
	fmt.Printf("折扣后的价格为：%.2f\n", Z) // %.2f 保留浮点数两位
}
