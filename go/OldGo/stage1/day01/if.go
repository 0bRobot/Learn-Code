package main

import "fmt"

func main() {
	//条件判断
	if1()
}

// 条件判断
func if1() {
	a := 3
	b := 5
	if a < b { //单 分支
		fmt.Println("a大于b", a, b)
	}

	if a >= 3 { //双分支 判断
		fmt.Println("条件1满足 a 大于或者等于", a) // 条件 为true  时执行
	} else {
		fmt.Println("条件1 不满足a为：", a) // 条件为 false 时执行

	}
	fmt.Println()
	if a > 4 { // 多分支    // 判断条件1  满足时执行
		fmt.Println("a大于1满足")
	} else if b < 4 { // 条件1 不满足时  判断条件2  是否满足
		fmt.Println("b小于5满足")
	} else { // 以上两个都不满足时 执行
		fmt.Printf("a的值为%v,b的值为%v\n", a, b)
	}

	if num := 10; num%2 == 0 { // num  变量只在 if 分支中有效  外部不可见
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	if score := 65; score >= 90 { // source 变量只在 if 分支中有效  外部不可见
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
