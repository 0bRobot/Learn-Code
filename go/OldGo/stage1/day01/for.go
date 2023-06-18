package main

import "fmt"

func main() {
	for1()
	for2()
	for3()
	for4()
}
func for1() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i)    变量 i 的作用域只在for循环中
}

func for2() {
	i := 0
	for i < 9 { //for循环的初始语句和结束语句都可以省略
		fmt.Println(i)
		i++ // 自增语句  条件满足退出循环
	}
	fmt.Println(i)
}

func for3() {
	i := 0
	for ; i <= 10; i++ { // for循环的初始语句可以被忽略，但是初始语句后的分号必须要写
		fmt.Println(i) //10
	}
	fmt.Println(i) // 11
}

func for4() {
	i := 1
	for {
		if i > 5 { // 判断 条件是否满足  满足则执行  break语句
			break //  退出当前 for 循环
		}
		i++ // i 自增
		fmt.Println("...")
	}
}
