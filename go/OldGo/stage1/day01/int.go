package main

import "fmt"

var (
	i1 = 0b1100  // 二进制
	i2 = 0o6743  //八进制进制  默认int 型
	i3 = 0xff98  // 十六进制
	i4 = 123_456 // _分隔让数字更直观
)

func main() {

	fmt.Println(i1, i2, i3, i4)

	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10    %d表示十进制
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77      %o 表示八进制

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff    %x 表示十六进制
	fmt.Printf("%X \n", c) // FF
}
