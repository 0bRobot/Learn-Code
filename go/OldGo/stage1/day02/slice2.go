package main

import "fmt"

func main() {
	//slice1()
	//slice2()
	//copyDemo()
	//forSlice()
	//appendSlice()
	//deleteSlice()
	lianxi()

}
func slice1() {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//默认切片的容量是从切片的开始索引到数组的最后
	// max : 指定了最大能切到的索引位置，影响切片后的容量
	B := A[0:3:6]       //  0 <=low<high<=max
	fmt.Println(B)      //[1 2 3]
	fmt.Println(len(B)) //3
	fmt.Println(cap(B)) //6

}

func slice2() {
	//make函数做初始化切片
	A := make([]int, 3, 8)

	var AA []int
	AA = make([]int, 4) // make函数做初始化切片   容量可以省略  但是长度和容量相等
	fmt.Println(A)      // [0 0 0]
	fmt.Println(len(A)) // 3
	fmt.Println(cap(A)) // 8
	fmt.Println()
	fmt.Println(len(AA), cap(AA), AA) // 4 4 [0 0 0 0]

	mun := 3
	imun := 10
	var CC []string
	CC = make([]string, mun, imun) //  make 函数中 元素个数和容量 可以是变量
	fmt.Println(CC)
	fmt.Println(len(CC), cap(CC))

	// 两个切片不能进行比较 但是 切片对应索引的值 可以进行比较
	fmt.Println(AA[0] == A[0]) // ture
	xx := make([]int, 1, 2)
	yy := make([]int, 1, 2)
	fmt.Println(xx[0] == yy[0]) //ture
	//fmt.Println(xx == yy)
	//var jj []int
	//var yy []int
	//fmt.Println(jj == yy)

}

func copyDemo() {
	E := []int{1, 2, 4, 5, 67, 8}
	F := make([]int, 3, 6) // 目标的F切片 长度是3 容量是6
	copy(F, E)             // 把切片E中的值 拷贝到切片F中。
	fmt.Println(F)         // [1 2 4]
	//F[5] = 40  //  这里编译 运行 会报panic错误     对应索引进行赋值  切片长度不存在。

	X := make([]int, len(E)) // 注意 切片拷贝时目标切片的长度 一定要等于大于原切片
	copy(X, E)               // 把切片E中的值拷贝到切片X中
	X[5] = 400
	fmt.Println(X) // [1 2 4 5 67 400]

}

func forSlice() {
	s := []int{1, 2, 3, 4, 56, 6, 8, 9}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}
	fmt.Println()
	s1 := []string{"hello", "test", "MrLin"}
	for _, v := range s1 {
		fmt.Printf("%s ", v)
	}

}

func appendSlice() {
	// append 函数往切片中追加元素
	// 使用append函数时必须使用变量接受返回值
	var s = []string{"测试", "北京"}
	fmt.Println("append前", len(s), cap(s)) //  2 2
	// 追加一个
	s = append(s, "上海")                    //append 函数可能触发切片的扩容  切片扩容之后会有新的底层数组， 需要更新新的变量s
	fmt.Println(s)                         // [测试 北京 上海]
	fmt.Println("append后", len(s), cap(s)) // 3 4

	// 追加多个
	x := []string{"成都", "四川"}
	x = append(x, "北京", "重庆", "天津", "上海")   //  追加多个
	x1 := append(x, "北京", "重庆", "天津", "上海") //  可以追加到新的切片中(未初始化) ，  并且有两次追加的内容
	fmt.Println(x)                          //             [成都 四川 北京 重庆 天津 上海]
	fmt.Println(x1)                         // [成都 四川 北京 重庆 天津 上海 北京 重庆 天津 上海]

	s2 := []string{"河北"}
	_ = append(s2, "上海", "北京")
	fmt.Println(s2)

}

func deleteSlice() {
	// 从切片中删除元素
	a := []int{1, 8, 2, 3, 4, 5}
	a = append(a[:1], a[2:]...) // 利用append特性  只要切片中 追加下标非1的值   就做到删除的对应的值
	// a = append(a[:inx],s[:inx+1]...)
	// 删除多个连续的值
	//删除切片 `a` 中从索引 `i` 至 `j` 位置的元素：`a = append(a[:i], a[j:]...)`
	fmt.Println(a)

	// 多维切片
	AA := [][]int{{1, 2, 3}, {4, 5, 6}, {5, 6, 7}}
	fmt.Println(AA)

	D := make([][]int, 3)
	fmt.Println(D)
}

func lianxi() {
	var a = make([]string, 5, 10) // make  初始化  5个空值    往后 继续追加
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)              //  [     0 1 2 3 4 5 6 7 8 9]     前面 5个空值
	fmt.Println(len(a), cap(a)) // 15   触发自动扩容 20  不确定

	var B = make([]string, 3, 8)
	fmt.Println(B)
	fmt.Printf("%#v", B) // []string{"", "", ""}

}
