package array

import (
	"fmt"
)

func Test1() {

	//	数组是同一种数据类型元素的集合。
	//	在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化  数组下标  从0 开始

	//数组的定义   var 关键字  数组名字 [元素数量]  类型
	var arr1 [5]int
	fmt.Println(arr1[1])
	arr2 := [...]string{}
	//使用 简短声明  定义 空数组

	fmt.Println(arr2)

	//数组的初始化   声明是初始化
	var arr3 [3]int // 默认该类型的0值

	fmt.Println(arr3[2])
	var arr4 = [4]string{"hello", "test1"}
	fmt.Println(arr4[1])
	var arr5 = [3]int{1, 2, 3}
	fmt.Println(arr5)

	//	初始化2
	var arr6 = [...]int{2, 4, 6, 7, 8, 0}
	fmt.Println(arr6)

	//指定索引来初始化

	var arr7 = [5]int{0: 1, 3: 4, 4: 5}
	fmt.Println(arr7[0], arr7[3])

}

func Test2() {

	//	数组的遍历
	arr1 := [...]string{"上海", "北京", "天津", "上海", "南京"}

	//fmt.Println(len(arr1))
	//求数组最大个数
	//	方法1  for 循环遍历
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	fmt.Println("=================分割线==========================")

	//	方法2  range
	for _, v := range arr1 { // _ 匿名变量
		//fmt.Println(i)
		fmt.Println(v)
	}
}
func Test3() {
	//	  多维数组
	//	多维数组的定义
	var marr [3][2]string
	fmt.Println(marr)
	marr1 := [3][3]string{{"东北", "黑龙江", "辽宁"}, {"陕西", "西安", "宝鸡"}, {"河北", "邯郸", "石家庄"}}
	//fmt.Println(marr1[1][1])

	//	多维数组的遍历
	for _, v := range marr1 { // 在外层循环中，使用range marr1来遍历一维数组的每个元素，其中_表示索引，v表示对应元素的值。对于这个示例，每个元素都是一个一维数组。
		//fmt.Println(i, v)
		for _, v := range v { // 在内层循环中，使用range v来遍历内层一维数组的每个元素。这里也使用了_表示索引，v表示对应元素的值。对于这个示例，每个元素都是一个字符串
			fmt.Println(v)
		}
	}
	// 二维数组 只有第一层 可以自动推倒数组个数
	marr2 := [...][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(marr2)
}
func Test4() {
	// 数组是值类型
	arr1 := [3]int{10, 20, 30}
	fmt.Println(arr1)
	fmt.Printf("%p\n", &arr1)
	arr1[0] = 40
	fmt.Println(arr1)
	fmt.Printf("%p\n", &arr1) // 数组是一个值类型，当对数组进行修改时，实际上是修改数组元素的值，而不是修改数组本身的地址。因此，数组的内存地址保持不变
}

//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值

func Test5() {
	arr1 := [3]int{1, 2, 3}
	modifyArr(arr1)
	fmt.Println(arr1)
}

func modifyArr(x [3]int) { // modifyArr函数接收到的是数组arr1的副本。当在modifyArr函数中修改副本的元素值时，不会影响到原始数组。
	x[0] = 11
	fmt.Println(x) // 拷贝值副本 已经被修改
	//要实现对数组的修改，可以通过传递数组的指针来实现
}

// 自动推导  ...

func Test6() {
	arr1 := [...]int{2, 3, 4, 5, 6}
	modifyArr1(arr1)
	fmt.Println("外部打印", arr1)
}
func modifyArr1(x [5]int) { // 函数接受数组 类型 不能使用 x [...]int  自动推导
	x[3] = 999
	fmt.Println("modifyArr函数内部打印", x)
}
