package main

import "fmt"

func main() {

	f12344 := adder()
	res1 := f12344()
	fmt.Println(res1) // 1
}

// 闭包
func adder() func() int {
	//  返回值为 函数
	var x int
	// 函数内部使用了它外部函数的变量X
	f := func() int {
		// 返回值 int
		x++
		return x
	}
	// 把匿名函数当成了返回值返回了
	return f
}
