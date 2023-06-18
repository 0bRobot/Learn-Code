package main

import (
	"fmt"
)

func main() {
	run01()
}

func run01() {
	//调用 fun02函数  传入三个实参  最后一个是函数类型
	//fun02(10, 30, add1)
}

//定义了 函数 函数的形式参数 有两个int 类型  和一个函数类型
//和下面fun02的方法 一样
type zz func(int, int) int // zz  为新类型
func fun02(x, y int, z zz) int { // z  为第三个形参   zz 为形参的类型
	res := z(x, y)
	fmt.Println(res)
	return res
}

//func fun01(x, y int, z func(int, int) int) int {
//
//}

func add1(x, y int) int {
	return x + y
}
func sub1(x, y int) int {
	return x - y
}

// 命名返回值
// 1. 函数内部声明了一个变量
// 2. 返回值是res

func f11(x, y int, s string) (res func(int, int) int) {
	// 声明一个函数 分别形参有 3个  返回值是一个函数
	switch s {
	case "+":
		return add1
	case "-":
		return sub1
	}
	return
}
func f12(x, y int, s string) func(int, int) int {

	var res func(int, int) int
	switch s {
	case "+":
		return add1
	case "-":
		return sub1
	}
	return res
}
