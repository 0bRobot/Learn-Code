package main

import (
	"fmt"
)

const pi float64 = 3.14          // 常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值
const pVersion = "linux  amd 64" // 类型可根据值自动推导   可省略
const (
	sysVersion string = "windows 10"
	sysOs      string = "10"
	abc        int    = 23
)
const (
	x1 = iota // 0                  const中每新增一行常量声明将使iota计数一次
	x2        // 1
	x3        // 2
)

func main() {
	const ( // const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。
		n1 = 100
		n2
		n3
	)

	const ( // const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。
		m1 = 10
		m2
		_
		m4
	)
	const (
		a1 = iota * 2
		a2
		a3
		a4
	)
	const (
		_  = iota             //利用iota 声明存储单位的常量
		KB = 1 << (10 * iota) //<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	var nn int64
	nn = 1 << 10 //  1012  左移两位 1024

	const (
		_  = iota
		ks = iota
		bs = iota
	)

	const (
		aa, bb = iota + 1, iota + 2
		cc, dd
		ee, ff
	)

	fmt.Println(pi)
	fmt.Println(pVersion)
	fmt.Println(sysVersion, sysOs, abc)
	fmt.Println(n1, n2, n3)
	fmt.Println()
	fmt.Println("iota", x1, x2, x3)
	fmt.Println("iotaSkip", m1, m2, m4)
	fmt.Println(nn)
}
