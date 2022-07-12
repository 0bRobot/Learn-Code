package main

import "fmt"

func main() {
	// 一维 数组的定义
	// 不初始化
	var arr1 [3]int    //定义数组 未初始化  有默认值   长度为 3
	var arr2 [4]string // 类型为字符串  未初始化  长度为 3
	var arr3 [4]bool   //  浮点型 [0 0 0 0]  bool类型  默认 为 [false false false false]

	fmt.Println(arr1) // 输出  [0 0 0]
	fmt.Println(arr2) // 输出 为 空 [   ]
	fmt.Println(arr3) // [0 0 0 0]
	//数组初始化
	var arr4 [5]int = [5]int{1, 7, 8, 4, 9} //完全初始化
	var arr5 [5]int = [5]int{1, 9}          // 不完全初始化
	var arr6 [5]int = [5]int{}              // 不初始化

	fmt.Println(arr4) // [1 7 8 4 9]
	fmt.Println(arr5) // [1 9 0 0 0]
	fmt.Println(arr6) //[0 0 0 0 0]

	// 数组定义的简洁语法
	// -----------------   注意简洁语法的{} 不能省略------------------
	arr7 := [6]int{1, 6, 7, 9, 8, 9} //   //完全初始化
	arr8 := [6]int{1, 9}             // 不完全初始化
	arr9 := [6]int{}                 // 不初始化
	fmt.Println(arr7)
	fmt.Println(arr8)
	fmt.Println(arr9)

	// 数组元素个数不确定 ？？ 如何定义
	//  ---====== 数组元素个数不确定  只能使用短语法定义   元素个数不确定 编译器 自动推导
	// 长语法 元素个数不确定  ---------------------------------
	var arr14 = []int{1, 2, 4, 6, 123, 131432, 23, 35, 24}
	fmt.Println(arr14)

	Arr10 := [...]int{1, 6, 8, 9} //元素个数不确定  //完全初始化
	Arr11 := [...]int{}           // 不初始化    // 不存在 不完全初始化

	fmt.Println(Arr10)      // [1 6 8 9]
	fmt.Println(len(Arr10)) //求数组元素个数
	fmt.Println(Arr11)      // []
	fmt.Println(len(Arr11))

	//多维数组的定义
	var Darr [2][3]int                              // 未初始化化
	var Darr1 [2][2]int = [2][2]int{{6}, {2}}       // 不完全初始化
	var Darr2 [2][2]int = [2][2]int{{6, 2}, {1, 2}} // 完全初始化

	fmt.Println(Darr)  //  [[0 0 0] [0 0 0]]
	fmt.Println(Darr1) // [[6 0] [2 0]]
	fmt.Println(Darr2) //[[6 2] [1 2]]

	// 多维数组 简介语法
	T1 := [3][4]int{}
	T2 := [3][5]int{{1, 4, 6, 78, 89}, {9, 6, 1}, {}}

	fmt.Println(T1) // [[0 0 0 0] [0 0 0 0] [0 0 0 0]]
	fmt.Println(T2) // [[1 4 6 78 89] [9 6 1 0 0] [0 0 0 0 0]]

}
