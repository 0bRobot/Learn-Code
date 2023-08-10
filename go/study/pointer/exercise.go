package pointer

import (
	"fmt"
)

// 编写一个函数，接受一个整数参数，将该参数的值加倍，并通过指针将结果返回。
func dome1(x *int) {
	*x *= 2 // *x = *x * 2
}

func DPointer2() {
	mun := 10
	dome1(&mun) // 函数传入一个指针
	fmt.Println(mun)

}

// 创建一个函数，接受一个整数切片的指针作为参数，将切片中的每个元素都翻倍。
func dome2(x *[]int) {
	for i := range *x {
		(*x)[i] = (*x)[i] * 2 // 简写为 (*s)[i] *= 2
		// 将切片中索引为 i 的元素翻倍，使用指针访问切片的元素
		// 使用 () 可以改变表达式的优先级。(*x)[i] 使用了括号来明确指定优先级，确保先对指针 x 进行解引用，然后再使用索引 i 来访问切片中的元素。
	}
}
func DPointer3() {
	s := []int{1, 2, 3, 4}
	dome2(&s)
	fmt.Println(s)
}

// 实现一个函数，接受两个整数指针作为参数，交换这两个指针所指向的整数的值。
func dome3(x, y *int) {
	*x, *y = *y, *x
}
func DPointer4() {
	a := 20
	b := 50
	fmt.Println("交换之前", a, b)
	dome3(&a, &b)
	fmt.Println("交换之后", a, b)
}

// 编写一个函数，接受一个整数切片的指针作为参数，返回切片中最大元素的值以及其索引。
func dome4(x *[]int) (mix, num int) {
	for n, v := range *x {
		if v > mix {
			mix = v
			num = n
		}
	}
	return mix, num
}

func DPointer5() {
	s := []int{3, 4, 6, 9, 54, 12, 9, 574, 98, 35, 13}
	mix := 0
	index := 0
	mix, index = dome4(&s)
	fmt.Println("切片中的最大值", mix, "\n对应的索引是", index)
}

// 创建一个函数，接受一个整数指针和一个整数切片的指针作为参数，将整数指针所指的值添加到切片中。
func dome5(x *int, y *[]int) {
	*y = append(*y, *x)
}
func DPointer6() {
	a := 45
	s := []int{6, 4, 3, 24, 34}
	dome5(&a, &s)
	fmt.Println(s)
}
