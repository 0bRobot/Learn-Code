package main

import (
	"fmt"
)

func main() {
	// 基本运算和的运算符
	//1. 运算符
	//2.表达式
	//3.双目运算符
	//4. 字面量
	// 5.表达式是可以求值的

	//拓展 变量的定义与常数的定义
	var a int = 4
	var b int = a * 4
	const c = 4
	var d = c * a * b
	//const e = c*a  常量的赋值  表达式中的值 不能为变量  可以为常量
	const e = c * c
	const f = c * e
	fmt.Printf("a = %d,b= %d\n", a, b)
	fmt.Println(d)
	fmt.Println(c, e, f)

	//算术运算符   +  -  *   /   %
	fmt.Println(5/2, 5.0/2.0) //变量的类型为整数  结果为整数
	//取余   类型只能是整数

	//关系运算符   == !=  <   >  <=  >=  双目运算符  需要两个表达式

	//逻辑运算符   &&   ||   !
	// 逻辑与 &&  两个同时为真 则为真     ||  逻辑或 一个为真 则为真   两个为假 则为假
	// ! 非运算符  相当于 取反     真  返回假   假返回真   非运算符  单目运算符
	fmt.Println(!(43 > 23))
	// 位运算符  & |  ^  <<   >>   按位与 &  按位或 |    按位异或 ^    左移运算符 <<   右移运算符 >>
	//位运算符 都是针对 二进制来运算的    <<  理解为 乘以2 的N次方    >>  除以2的N次方
	o := 20
	k := 1 << 10
	println(o << 2)  //乘2的2 次方
	println(o>>2, k) // 除以2 的2 次方
	//赋值运算符 = += -= /= *=  %=   <<+  >>=  &=  ^= |=

	var x int = 11
	x += 2
	fmt.Printf("%d\n", x)

	var UU int = 20
	var pp int
	pp += UU
	//pp = pp + UU
	fmt.Println(pp)

}
