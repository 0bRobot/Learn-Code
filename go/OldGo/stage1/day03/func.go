package main

import "fmt"

func main() {
	sum1 := addNum(8, 8) // 函数的调用
	fmt.Println(sum1)
	//fmt.Println(addNum(5, 7))
	//多个返回值
	a, b := f1(6, 8)
	fmt.Println(a, b)

	AA := addSome(4, 6, 7, 1, 3)
	fmt.Println(AA)
	fmt.Println(addSum1(3, 1, 2, 3, 4, 5, 6))
	//fmt.Println(reMap())

	soFunc("")
}

// 函数的定义
func addNum(x, y int) int {
	// 有多个返回值用()
	//  有返回值必须使用 return
	// 函数的返回值 也需要说明类型  定义
	return x + y

}
func f1(x int, y int) (int, int) {
	// 有多个返回值
	aa := x + x
	bb := y + y
	return aa, bb

}

// 求一堆int的和  结束数量不确定
// 可变参数 就是切片
func addSome(x ...int) int { //接收数量不确定
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func addSum1(x int, y ...int) int {
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

func reName(x, y int) (AA int, BB int) {
	//命名返回值
	//相当于在函数内部声明了返回值变量
	// AA := x + x
	// BB := y + y  // 就不需要 进行简短声明
	AA = x + x
	BB = y + y
	return AA, BB
}

func reMap() (mm map[int]string) {
	// 函数的返回值类似是map
	// 命名返回值需要用括号括起来
	// 命名返回值 等同于使用var 声明了一个map  map 的使用必须初始化
	// 初始化 mm map
	mm = make(map[int]string, 4) //未初始化 运行报错panic: assignment to entry in nil map
	mm[1] = "北京"
	mm[2] = "上海"
	return mm
}

func soFunc(x string) []int {

	if x == "" {
		// 如果函数传入值为空
		return nil //  没必要返回[]int{ (长度为0的切片)   返回一个nil 即可
	}
	// 如果传入参数 返回对的元素
	return []int{1, 2, 3, 4}
}
