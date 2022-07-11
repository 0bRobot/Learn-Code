package main

import "fmt"

func userp() {
	fmt.Println("-----------------------------------------")
	fmt.Println("自定义函数输出，无返回值")
}

func Add(a, b int) (r int) { //int  返回值 名称和 类型     形参
	return a + b // 函数有返回值    return 不能省略
}
func Add2(a, b int) int { //int  返回值 名称和 类型     //  返回值 名称 可省略
	return a + b // 函数有返回值    return 不能省略
}

// 一个函数实现加减乘除
// 传入+ 好  计算加法  传入 -  计算减法  *  /

func op(s string) func(int, int) int {
	switch {
	case s == "+": //op函数传入对应的符号 做对应的计算
		return func(a, b int) int { //return  匿名函数的值
			return a + b
		}
	case s == "-":
		return func(a, b int) int {
			return a - b //return  匿名函数的值
		}
	default:
		panic("操作不支持") //go 语言提供的内置函数  提示反生了错误

	}
}

//函数的可变参数

func test(AAA ...int) { //函数接收不定数量的参数  类型需确定
	fmt.Println(AAA)
	test2(AAA...) //可变参数的嵌套  不能直接使用 变量名  传递的是这个参数的切片
}
func test2(AAA ...int) { //函数接收不定数量的参数  类型需确定

}

func main() {
	fmt.Println("函数")

	// func 关键字  函数的名称 (函数参数的列表)  函数返回值列表 { //函数体 }
	//函数的调用
	F := Add(3, 5) //函数的参数  实参  //函数有返回值    变量F 接收函数的返回值
	_ = Add(3, 5)  //函数的参数  实参  //函数有返回值    变量F 接收函数的返回值  //占位符 忽略不适应的函数
	fmt.Println(F)
	E := Add(6, 8)
	fmt.Println(E)
	userp() // 无返回值得调用

	// 匿名函数   定义的时候函数没有名字  把函数的结果赋值给变量   且匿名函数 只能在主函数体内
	GG := func(a int, b int) int {
		return a + b
	}
	fmt.Println(GG(1, 2)) //输出函数结果

	// 匿名函数 2
	println(func(a int, b int) int {
		return a + b
	}(5, 4)) // 直接输出
	//func(a int, b int) int {
	//	return a + b
	//}(5, 4)

	// 加减法函数的调用
	opFUNC := op("+")                  //第一个函数的返回给变量opFUNC
	fmt.Println("匿名函数+", opFUNC(9, 6)) // 传入匿名函数的值
	opFUNC = op("-")                   //第一个函数的返回给变量opFUNC
	fmt.Println("匿名函数-", opFUNC(6, 9))
	// 传入匿名函数的值

	Y1, Y2 := 5, 9
	R1, R2 := swap(Y1, Y2)
	fmt.Println(R1, R2)
	fmt.Println(Y1, Y2) // 原来的值 未被修改  值传递  也就传值

	// 怎么实现传址呢
	//闭包
	func() {
		Y1, Y2 = Y2, Y1 // 从外层获取的  函数未定义
	}()
	fmt.Println("y1,y2 闭包", Y1, Y2) // 输出发现原始值已经改变

	test(12, 234, 345, 54, 677, 23, 234) // 函数的可变参数 实参输入个数不确定

	// 调用输出为数组
}

//  值传递 	值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。   也就数传值
//  引用传递 	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。   传址
// go 语言函数默认  传值

func swap(a int, b int) (int, int) { // 返回两个数  定义两个返回类型
	return b, a // 交换返回
}

//func function_name( [parameter list] ) [return_types] {
//函数体
//}
//func：函数由 func 开始声明
//function_name：函数名称，参数列表和返回值类型构成了函数签名。
//parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
//return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
//函数体：函数定义的代码集合
