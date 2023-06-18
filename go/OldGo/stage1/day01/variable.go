package main

import "fmt"

// 标准变量的声明
var name string = "张小明"
var gender string // 变量声明未赋值

var ss = "测试" // 函数外部声明的变量  不使用  不影响运行

// 批量变量声明
var (
	age         = 20 // 根据值自动推导 数据类型
	addr string = "北京市朝阳区"
)

func main() {
	gender = "男"           // 变量赋值  赋值 赋对应的正确的值
	phone := "18199999999" // 简短变量声明 只能在函数内部

	//一次初始化多个变量
	var name1, age1, addr1 = "张小小", 19, "北京市海定区"

	// go语言中 变量声明必须使用 否则编译不过

	fmt.Println(name, gender, age, addr, phone)
	fmt.Println(name1, age1, addr1)

	// 单引号表示字符  双引号表示字符串

	var x = "a" //  字符串
	var s = 'a' //   字符
	fmt.Println(x, s)

}
