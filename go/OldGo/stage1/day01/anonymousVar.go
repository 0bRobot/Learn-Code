package main

import "fmt"

func foo() (bool, string) { //自定义函数  返回两个值
	return false, "error"

}

func main() {
	x, _ := foo() //  函数return 的 string值 赋值给匿名变量  接收不想要的值
	_, y := foo() // 匿名变量不占用内存空间，不会分配内  所以匿名变量之间不存在重复声明
	fmt.Println("x", x)
	fmt.Println("y", y)
}
