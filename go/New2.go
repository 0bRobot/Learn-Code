package main

import (
	"fmt"
	"math"
)

func main() {

	//1.2 基本的数据类型
	//bool  布尔
	var a bool = false
	var b string = "张三"
	var c int = 23
	var d uint = 123 // 无符号不能为负数
	var e float32 = 3.1415
	var f complex64 = 3234 + 23i

	//string  字符串
	// (u)int  (u)int8  (u)int16  (u)int32  (u)int64   整数 (u 无符号  有符号)
	//uintprt 指针
	//float32   float64 浮点数
	//complex64   complex128   复数
	fmt.Println(a, b, c, d, e, f)
	//类型转换
	// go语言无隐式类型装换   需要强制类型装换
	var ttt int32 = 100
	var yyy int64
	// ttt = yyy  不能把 int32变量的值赋值给int64
	//需要类型强制装换
	yyy = int64(ttt)

	var a2 int16 = 23
	var b1 float32
	b1 = float32(float64(a2)) // 多层装换
	fmt.Println(yyy)
	fmt.Println(b1)

	var ab, bc = 3, 4
	cc := math.Sqrt(float64(ab*ab + bc*bc)) //Sqrt 函数要求数据类型为float64
	println("第三条边的长度为:", cc)
	// 浮点数 使用科学计数法  使用IEEE754 标准

}
