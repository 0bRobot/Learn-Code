package main

import (
	"fmt"
	"math"
)

func main() {
	// 常量的定义
	// 使用const关键字
	//  单个常量的定义
	const Pi = 3.1415
	const r = 5
	const OsVersion = "2.3.1"

	// 多个常量的定义
	const a, b = 1, 2 // 常量未使用 可以编译通过
	const (
		c = 2
		d = 3
		f = 3.14535
	)
	//计算三角函数
	e := math.Sqrt(c*c + d*d) //常量自适应数据类型  不需要强制装换  如果 const 指定类型需要强制类型转换
	fmt.Println(e)

	fmt.Println("半径为是: ", r, " 圆周长为: ", Pi*r)
	fmt.Println("系统版本：", OsVersion)

	//枚举  enum
	// 变量首字母大写 默认public其他包可见
	const (
		AA = "NIhao"
		BB = "GO 语言"
		CC = "C 语言"
	)
	fmt.Println(AA, BB, CC)

	// 自增类型的枚举
	// iota  初始化为0  每次自增1
	const (
		DD = iota + 1 // 改变iota的初始值
		EE
		_ //跳过计数  占位符
		FF
	)
	const (
		ZZ = (iota + 1) * 2 // 改变iota的初始值  设置 步长为2
		XX
		YY
	)
	fmt.Println(DD, EE, FF)
	fmt.Println(ZZ, XX, YY)
}
