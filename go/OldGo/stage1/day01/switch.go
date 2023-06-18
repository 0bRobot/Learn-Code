package main

import "fmt"

func main() {
	//switch 语句可方便地对大量的值进行条件判断
	switch1()
	//一个分支可以有多个值，多个case值中间使用英文逗号分隔**
	switch2()
	// 分支表达式
	switch3()
}
func switch1() {
	a := 1

	switch a {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println(a)

	}
}

func switch2() {
	switch n := 6; n { // 变量在switch 内部
	case 1, 3, 5, 7, 9: // 多个case值中间使用英文逗号分隔
		fmt.Println("输入的是奇数")
	case 2, 4, 6, 8:
		fmt.Println("输入的是偶数")
	default:
		fmt.Println("输入错误")
	}

}

func switch3() {
	age := 48
	switch {
	case age <= 24:
		fmt.Println("你真年轻")
	case age > 25 && age < 30:
		fmt.Println("真是努力的时候")
	case age > 30 && age < 39:
		fmt.Println("你结婚生娃了没有")
	case age > 40 && age < 70:
		fmt.Println("凑合过吧")
	case age >= 70 && age <= 90:
		fmt.Println("颐养天年")
	case age >= 100:
		fmt.Println("老神仙")
	default:
		fmt.Println("老妖怪")

	}

}
