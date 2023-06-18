package main

import "fmt"

func main() {
	//
	//defer1()

	fmt.Println(f21())
	fmt.Println(f22())
	fmt.Println(f23())
	fmt.Println(f24())
}

func f21() int { //5
	// defer语句执行的时机就在返回值赋值操作后，RET指令执行前
	x := 5
	defer func() {
		x++
	}()
	// 返回值 x 等于5  返回值赋值已经结束   然后在执行 匿名函数 x++  x赋值为6
	// 只改了x的值 没有更改 返回值 return 的值
	return x
}

func f22() (x int) { // 命名返回值 返回x 的值  6
	defer func() {
		x++
	}()
	return 5 // 返回值x == 5    defer 操作 x++ x 等于6
}

func f23() (y int) { // 命名返回值 y
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f24() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
