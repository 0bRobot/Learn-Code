package main

import "fmt"

func main() {

}

//  匿名函数
// 函数内部不能在定义函数

func ff12() func() { // 主函数返回一个函数类型
	//go 语言不能在函数内  再使用 func 关键字 在定义函数
	f123 := func() {
		// 匿名函数
		// 匿名函数因为没有函数名，所以不能直接调用  需要保存到某个变量或者作为立即执行函数
		fmt.Println("匿名函数")
	}
	return f123
	//var fff1 func()=func(){
	//	fmt.Println("匿名函数")
	//}
}
func ff13() {

	func() {
		fmt.Println("匿名函数立即执行")
	}()
	// 匿名函数因为没有函数名，不能直接调用  需要保存到某个变量或者作为立即执行函数
	return

}
