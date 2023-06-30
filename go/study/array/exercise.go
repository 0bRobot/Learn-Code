package array

import (
	"fmt"
	"sort"
)

// 编写一个函数，接受一个整数数组作为参数，返回数组中的最大值和最小值

func ExecCode1() {
	arrEx := [...]int{3, 5, 7, 9, 1, 0, 4}
	fmt.Println("该数组的值是", arrEx)

	min, max := returnMun(arrEx) // 接受返回值给变量 min max
	fmt.Println("最大值", max)
	fmt.Println("最小值", min)

}
func returnMun(x [7]int) (int, int) { //
	if len(x) == 0 {
		return 0, 0
	}
	min := x[0] // 假设第一个元素为最小值
	max := x[0] // 假设第一个元素为最大值
	for _, v := range x {
		if v < min {
			min = v // 更新最小值
		}
		if v > max {
			max = v // 更新最大值
		}
	}
	return min, max
}

//编写一个函数，接受一个整数数组作为参数，返回数组中所有奇数的和。

func ExecCode2() {
	munArr := [...]int{2, 3, 4, 5, 6, 7, 8, 3, 11}
	fmt.Println("该数组是", munArr)
	sum := sumMun(munArr)
	fmt.Println("数组中奇数的和是", sum)
}
func sumMun(x [9]int) int {
	sum := 0
	for _, v := range x {
		if v%2 != 0 {
			sum = sum + v
		}
	}
	return sum
}

//编写一个函数，接受一个整数数组和一个目标值作为参数，返回数组中两个元素的索引，使得它们的和等于目标值。假设数组中只有一个解。

func ExecCode3() {
	T := 5
	arr1 := [...]int{2, 3, 4, 56, 84, 2, 5, 13, 34} // 一个解
	index, index1 := sumT1(arr1, T)
	fmt.Println("对应的索引为", index, index1)
	// 只输出 一个解 0  5
}
func sumT1(x [9]int, T int) (index1, index2 int) {
	for i := range x {
		for j := range x {
			if i+j == T {
				return i, j

			}
		}
	}

	return 0, 0
}

func ExecCode4() {
	T := 5
	arr1 := [...]int{2, 3, 4, 56, 84, 2, 5, 13, 34} // 。有多个解
	sumT(arr1, T)

	// 只输出 一个解 0  5
}
func sumT(x [9]int, T int) {
	for i := range x {
		for j := range x {
			if i+j == T {
				fmt.Println("对应的索引下标是：", i, j)

			}
		}
	}

}

//  编写一个函数，接受一个整数数组作为参数，将数组中的元素按照升序排序，并返回排序后的数组。

func ExecCode5() {
	arr1 := [5]int{7, 9, 3, 5, 2}
	fmt.Println(SortArr1(arr1))

}
func SortArr1(x [5]int) [5]int {
	arr := x          // 将原始数组 x 复制给 arr
	sort.Ints(arr[:]) // 使用 sort.Ints 函数对切片 arr 进行排序
	return arr        // 返回排序后的数组 arr
}

//可以直接对原始数组进行排序，可以避免内存复制，提高性能。
//使用指针传递：可以将原始数组的指针传递给排序函数，以便直接在原始数组上进行排序。这样可以减少对数组的复制，并且函数内部的排序操作将直接影响原始数组。

func ExecCode6() {
	arr1 := [5]int{7, 9, 3, 5, 2}
	SortArr2(&arr1) // 将数组的地址传递给排序函数
	fmt.Println(arr1)
}

func SortArr2(x *[5]int) {
	sort.Ints(x[:]) // 直接对原始数组进行排序
}

//  编写一个函数，接受一个整数数组作为参数，返回数组中所有元素的平均值。

func ExecCode7() {
	arr1 := [8]int{3, 5, 6, 7, 8, 2, 7, 9}
	ave := averageV(arr1)
	fmt.Println("平局值是", ave)
}
func averageV(x [8]int) float64 {
	sum := 0.0
	for _, v := range x {
		sum += float64(v)
	}

	return sum / float64(len(x))
}
