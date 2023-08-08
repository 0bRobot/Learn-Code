package funcDome

import (
	"fmt"
	"math"
)

// 1.编写一个函数 counter，它返回一个闭包。每次调用闭包时，返回的值递增 1。

func counter() func(int) int {
	var x int // 闭包是一个函数，它可以访问其词法作用域之外的变量。换句话说，闭包可以“捕获”并记住在其定义时可见的变量。
	return func(i int) int {
		x++
		return x
	}
}
func AdvFunc01() {
	var f = counter()
	fmt.Println(f(10))
	fmt.Println(f(20))
}

// 2.创建一个函数 calculatePower，接受一个浮点数参数 base，并返回一个匿名函数，该匿名函数接受一个整数参数 exponent，返回 base 的指数幂结果。
func calculatePower(base float64) func(int) float64 {
	return func(i int) float64 {
		return math.Pow(base, float64(i)) //  返回 base 的 i 次幂结果
	}
}

func AdvFunc02() {
	f := 3.0
	// 调用 calculatePower 函数，传入浮点数 f，并返回一个计算幂次方的匿名函数
	powerFun := calculatePower(f)
	// 调用 powerFun，传入整数参数 3，计算 3.0 的 3 次幂
	fmt.Println(powerFun(3))
}

// 3.实现一个函数 makeAdder，接受一个整数参数 x，返回一个接受一个整数参数 y 的匿名函数，返回 x + y 的结果。

// makeAdder 接受一个整数参数 x，返回一个匿名函数
// 这个匿名函数接受一个整数参数 y，返回 x + y 的结果
func mackAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func AdvFunc03() {
	adder := mackAdder(10) // 创建一个匿名函数，将 x 设置为 10
	result := adder(5)     // 调用匿名函数，将 y 设置为 5，得到结果
	fmt.Println(result)

}

// 4.编写一个函数 filter，接受一个整数切片和一个判断函数作为参数，返回一个新的切片，其中包含原切片中所有满足判断函数条件的元素。

// filter 函数接受一个整数切片 s 和一个判断条件函数 x
// 返回一个新切片，其中包含满足条件 x 的整数。
func filter(s []int, x func(int) bool) []int {
	tmp := make([]int, 0)
	for _, v := range s { // 遍历输入的整数切片 s 中的每个元素
		if x(v) { // 调用判断条件函数 x，判断整数 v 是否满足条件
			tmp = append(tmp, v) // 如果满足条件，则将整数 v 添加到 tmp 切片中
		}
	}
	// 返回存储满足条件的整数的切片 tmp
	return tmp
}
func pD(n int) bool { // pD 函数接受一个整数 n，判断 n 是否为偶数，返回布尔值。
	return n%2 == 0
}
func AdvFunc04() {
	s := []int{1, 2, 3, 4, 5, 5, 6, 7, 3, 8, 67, 23, 44}
	// 调用 filter 函数，传入整数切片 s 和判断条件函数 pD，
	// 返回满足条件的整数切片 s1
	s1 := filter(s, pD)
	fmt.Println(s1)
}

// 5.创建一个函数 sequence，它返回一个闭包，每次调用闭包时返回一个递增的整数序列。
func sequence() func() int { // sequence 函数返回一个闭包，每次调用闭包时返回一个递增的整数序列
	s := 0              // 定义一个整数变量，用于记录当前的序列值
	return func() int { // 创建并返回一个闭包函数
		s++
		return s
	}
}

func AdvFunc05() {
	netsT := sequence() // 调用 sequence 函数，获取一个返回递增整数序列的闭包
	// 调用闭包函数，每次调用都返回一个递增的整数
	fmt.Println(netsT())
	fmt.Println(netsT())
	fmt.Println(netsT())

}
