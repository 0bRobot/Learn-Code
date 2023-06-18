package main

import (
	"fmt"
)

func main() {

	//arr1()
	//arr2()
	arr3()
}

func arr1() {
	// 数组是同一种数据类型元素的集合。
	//	固定长度，数组声明后长度便不能再修改  数组的长度 必须是常量
	//只能存储一种特定类型元素的序列
	const a int = 3
	var ar1 [10]string  //  数组的定义
	var ar2 [10 + 1]int //  表达式
	var ar3 [a]float64  //  a 为定义常量    // 索引范围 0  1   2  依次类推
	var ar4 [2]bool     // bool 类型的数组 默认为 false

	fmt.Println(ar1, ar2, ar3, ar4) //  未初始化的  int数组  默认值为0   string类型的数组为默认为空

	//	数组的初始化
	ar1[0] = "张三"
	ar2[9] = 34
	ar3[2] = 3.1415
	ar4[1] = true
	//数组支持 索引访问   索引的合法范围  0-len(array) -l  .  第一个下标为0
	// 数组索引下标 注意不要越界
	fmt.Println(ar1, ar2, ar3, ar4)

	sizeArray := len(ar1) - 1 // ar1 数组的下标最大值

	fmt.Println(sizeArray)

}

func arr2() {
	//使用初始化列表来设置数组元素的值
	var domeArray [3]int                         //数组会初始化为int类型的零值
	var dome1Array = [4]int{2, 3, 4, 5}          //使用指定的初始值完成初始化
	var dome2Array = [3]string{"北京", "上海", "天津"} //使用指定的初始值完成初始化

	fmt.Println(domeArray, dome1Array, dome2Array)

	var dome3Array [3]int
	var dome4Array = [...]int{1, 2, 3, 4}
	var dome5Array = [...]string{"张三", "Peter", "Tom", ""}
	fmt.Println(dome3Array, dome4Array, dome5Array)

	var dome6Array = [...]int{0: 123, 3: 139, 5: 554}
	var dome7Array = [...]string{1: "王老五", 3: "王麻子"}
	fmt.Println(dome6Array)
	fmt.Println(dome7Array, len(dome7Array))

}

func arr3() {

	var x1 = [...]int{2, 4, 5, 6, 7, 6: 4}
	//fmt.Print(x1)

	// for 循环遍历数组
	for i := 0; i < len(x1); i++ {
		fmt.Print(x1[i])
	}
	fmt.Printf("\n")

	// for range  遍历
	var x2 = [...]string{"张三", "王五", "赵柳", "王强"}
	for _, v := range x2 { // 匿名变量 丢弃索引
		fmt.Println(v)
	}

}
