package main

import "fmt"

func main() {
	//type1()
	type2()
}

func type1() {
	// 自定义函数类型  type 关键字
	type newFunc func(int, int) int
	// 声明一个变量 类型是自定义的函数类型
	var x newFunc
	// 给x新类型 赋值
	x = func(x int, y int) int {
		return x + y
	}
	//输出x 时传入对应参数
	fmt.Println(x(2, 4))
}

func type2() {
	//  函数签名  理解为  函数定义(声明)的格式
	// 函数的参数,返回值的类型和个数、顺序都要一样
	// 如接受两个int 类型 返回一个int类型
	type num123 func(int, int) int
	var xx num123
	xx = func(i int, i2 int) int {
		return i + i2
	}
	fmt.Println(xx(33, 34))
	var xx3 num123
	xx3 = func(i int, i2 int) int {
		return i - i2
	}
	fmt.Println(xx3(34, 54))

}
