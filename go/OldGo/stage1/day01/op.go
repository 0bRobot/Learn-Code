package main

import "fmt"

func main() {

	//  自增 自减  都是语句 不是运算符

	w11 := 3
	w11++
	//  w11 = w11++  这么写错误
	fmt.Println(w11)

	//  算术运算
	op()
	//关系运算符
	ver()

	//	逻辑运算符
	logic()
	// 位运算
	wei()
	// 赋值运算符
	values()
}

//  算术运算

func op() {
	a := 12
	b := 30
	c := a + b
	d := a % b

	fmt.Println(c)
	fmt.Println(d)

}
func ver() {
	a := "a"
	b := "a"
	d := 10.1
	c := 10.0
	fmt.Println(a == b)
	fmt.Println(c == d)

}

func logic() {
	aa := true
	bb := false
	cc := true
	fmt.Println(aa && bb) //  逻辑与  两者为 true 则为true   否则为false
	fmt.Println(aa && cc)
	fmt.Println(aa || cc) // 一个为true  则为true
	fmt.Println(!aa)      // 取反  aa 为false
	fmt.Println(!bb)
}
func wei() {
	a := 0b11011001
	b := 0b11111000
	c := 2
	fmt.Printf("%b\n", a&b) //  %b  输出二进制  二进制两位均为1才为1
	fmt.Printf("%b\n", a|b) //  二进制 两位有一个为1就为1
	fmt.Printf("%b\n", a^b) //  二进制 两位不一样则为1
	d := c << 2             // 二进制位 小数点 向左移动两位  低位补0  **左移n位就是乘以2的n次方
	fmt.Println(d)
	e := c >> 2
	fmt.Println(e) //**右移n位就是除以2的n次方    二进制位 小数点 向右移动两位
}

func values() {

	var a int64
	var b int
	a = 9 //简单的赋值运算符，将一个表达式的值赋给一个左值
	b = 3
	fmt.Println(a)
	a += a //   等价a = a+a
	fmt.Println(a)
	b *= b //  等价b=b*b
	fmt.Println(b)

}
