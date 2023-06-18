package main

import "fmt"

func main() {

	range1()
}

func range1() {
	i := "Hello go"
	for k, v := range i { //  for range  两个返回值  第一个返回值 k 是索引  第二个是在该索引位置的值  他们都是仅在 `for` 循环内部可见的局部变量
		fmt.Printf("%v,%c\n", k, v) //  %c  打印字符
	}
}
