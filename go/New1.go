// 定义文件所在包的名字  main  可执行文件
package main

import "fmt"

var x = 5
var (
	y = 5
	z = "ni hoa a"
)

// test := 345

// main 函数是main包下的特殊函数，代表可执行程序的入口
func main() {
	// 1. 单个变量的定义
	// 1.1 基础定义   var[关键字]   变量_name  变量的数据类型 = 初始值
	var name_1 string = "李四"
	// 1.2    var[关键字]   变量_name  = 初始值
	var age = 32
	// 1.2   var[关键字]   变量_name
	var age1 int      //整形默认初始值为0
	var name_2 string //  字符串为空

	// 多个变量的定义
	// 一行定义多个变量
	var ageN, ageN1 int = 23, 54
	var sex1, sex2 = '男', '女' //自动识别变量
	var birth1, birth2 = "1999-03-21", "1923-83-2"
	fmt.Println(ageN, ageN1, sex1, sex2, birth1, birth2)

	var ( // 定义多个变量
		a      = 5
		b      = "测试"
		c bool = false
	)
	// 简洁语法   // 只能应该到到函数体内部
	name3 := "王五"
	age4 := 43
	birth3 := "2002-23-23"

	fmt.Printf("a是%T b是%T,c是%T\n", a, b, c)

	fmt.Println(name_1, age, age1, name_2)
	fmt.Printf("%s\n", name_2)
	fmt.Println(name3, age4, birth3)
	fmt.Printf("%T,%T,%T\n", name3, age4, birth3)
	fmt.Println(x, y, z)

}
