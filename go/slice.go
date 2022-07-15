package main

import "fmt"

func main() {
	// 切片的定义
	var arr1 [6]int = [6]int{1, 2, 3, 4, 5, 6} //数组的定义
	// 数组下标    0 ,1 ,2 ,3, 4, 6
	var _ []int // 切片定义  ,slice

	fmt.Println("[1:3]", arr1[1:3]) // 输出 [2 3]   从下标为1开始取值  不包含下标1  取到下标 3 不包含下标3  半开半闭区间 不包含开始结束
	fmt.Println("[:3]", arr1[:3])   // 输出 [1 2 3]  结束闭区间 不包含 下标3
	fmt.Println("[1:]", arr1[1:])   // 输出 [2 3 4 5 6]  开始开区间 包含下标1
	fmt.Println("[:]", arr1[:])     // 输出 [1 2 3 4 5 6]  取所有

	s1 := arr1[1:3]
	s2 := s1[2:5]

	fmt.Println("s1[1:3]", s1)     //  [2 3]
	fmt.Println("s1- s2[2:5]", s2) //   [4 5 6]  s1 没有 4 5 6
	//切片的长度
	fmt.Println("s1 长度为", len(s1)) // 2
	fmt.Println("s1 长度为", len(s2)) // 3   s2 的4 5 6  来自哪里

	//切片的容量    cap() go 内置函数 求切片的容量

	fmt.Println("s1 容量为", cap(s1)) // 容量为 5  从  2, 3, 4, 5, 6 一直观察到结束
	fmt.Println("s1 容量为", cap(s2)) //  容量为 3
	//切片只对数组的局部或者全部的一种引用形式
	// 容量是切片 从索引开始到索引结束的

	// 切片的操作

	// 切片 追加元素  内置 append 函数  多接收多个值
	G := [...]int{1, 2, 4, 5, 6, 7, 12}
	DD := G[:4]
	fmt.Println(DD)
	DD = append(DD, 101) // 第一个追加
	// append 函数
	//fmt.Println(DD1, cap(DD1), cap(DD))
	fmt.Println(G)
	DD = append(DD, 8, 45) // 第二次追加    超出原数组元素个数 追加不生效
	//写法二
	DD = append(DD, []int{9, 66}...) // 类型  值  ... 展开操作符  展开切片 []int{123, 234}...
	//切片追加内容 超过数组内容   切片超过数组长度的时候  系统会创建一个新数组
	fmt.Println(G)

	//切片的复制 copy
	H := []int{1, 3, 5, 7, 9, 4}
	h1 := H[1:3]
	h2 := H[3:]
	fmt.Println(H)
	fmt.Println(h1)
	fmt.Println(h2)

	copy(h2, h1) // h1的值   复制到h2
	fmt.Println(H)
	fmt.Println(h1)
	fmt.Println(h2)

	fmt.Println("--------------------------------------")
	I := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("copy  前")
	i1 := I[0:4]
	i2 := I[5:7]
	fmt.Println(I)
	fmt.Println(i1)
	fmt.Println(i2)
	copy(i2, i1) //  后面的值复制到前面   只能复制 切片

	fmt.Println("copy  后 -------------------")
	fmt.Println(I)
	fmt.Println(i1)
	fmt.Println(i2)

	// 切片的初始化 二 make()

	TT := make([]int, 3, 6)  // 定义切片   make函数 （类型,长度，容量[可忽略]）
	TT1 := make([]string, 2) // 忽略容量
	TT2 := make([]int, 3)

	fmt.Println(TT, cap(TT), len(TT))    // [0 0 0]
	fmt.Println(TT1, cap(TT1), len(TT1)) // [ ]
	fmt.Println(TT2, cap(TT2), len(TT2))
	fmt.Printf("TT类型%T\nTT1的类型%T", TT, TT1)

	//make（)创建数组赋值
	//方法一
	t2 := make([]int, 3)
	t3 := []int{1, 23, 3}
	copy(t2, t3)
	fmt.Println(t2)
	t2 = append(t2, 105)
	fmt.Println(t2, cap(t2))
	//方法二
	s := make([]int, 3)
	a2 := s
	a2[0] = 1
	a2[1] = 2
	a2[2] = 3

}
