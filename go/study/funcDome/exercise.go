package funcDome

import (
	"fmt"
	"strings"
)

//1. 编写一个函数，接受一个整数切片作为参数，返回一个新切片，其中只包含偶数元素。

func even(x []int) []int {
	tmp := make([]int, 0, len(x))
	for _, v := range x {
		if v%2 == 0 {
			tmp = append(tmp, v)
		}
	}
	return tmp
}
func DomeFun01() {
	s := []int{1, 2, 3, 4, 5, 6, 8, 24, 23, 56, 768}
	fmt.Println((even(s)))
}

// 2. 创建一个函数，接受一个字符串作为参数，判断该字符串是否为回文（正序和倒序相同）
func strFunc(s string) bool {
	s = strings.ToLower(s) //  	切换为小写  以便不区分大小写
	lengs := len(s)
	// 使用循环判断字符串的前半部分和后半部分是否相等

	for i := 0; i < lengs/2; i++ {
		if s[i] != s[lengs-1-i] { // s[i] 表示字符串 s 中的第 i 个字符  际上是在比较字符串 s 中索引为 i 和索引为 lengs-1-i 处的两个字节（即字符）的值，这相当于比较这两个字符的ASCII码值或Unicode码点。
			return false
		}
	}
	return true
}

func DomeFun02() {
	fmt.Println(strFunc("racecar"))
	fmt.Println(strFunc("deed"))

	ff := "nihao"
	fmt.Println(ff[0]) // 110
}

//3. 实现一个函数，接受一个整数参数 n，返回一个 n x n 的二维切片，其中每个元素为其所在的行号和列号的和。

func slice01(n int) [][]int {
	x := make([][]int, n) // 创建一个 n x n 的二维切片

	for i := 0; i < n; i++ { // 填充二维切片中的每一行
		row := make([]int, n) // // 创建一个长度为 n 的整数切片，表示当前行
		// 对于每一行，我们创建一个长度为 n 的整数切片 row，然后在内部的循环中，为每个元素赋值为其行号和列号的和。最后，将每一行的切片 row 添加到二维切片 x 中，最终返回结果。
		for j := 0; j < n; j++ { // 填充当前行的每个元素
			row[j] = i + 1 + j + 1 // 元素的值为行号 i 和列号 j 的和
		}
		x[i] = row // 将填充好的行添加到二维切片 x 中
	}
	return x
}
func DomeFun03() {
	fmt.Println(slice01(4))
}

//4. 编写一个函数，接受一个字符串切片和一个字符串作为参数，返回切片中所有包含该字符串的元素

func strSlice(s []string, x string) []string {

	tmp := make([]string, 0) // 创建一个空切片 tmp，用于存储匹配的元素
	for _, v := range s {    // 遍历输入的字符串切片 s 中的每个元素
		if v == x {
			tmp = append(tmp, v) // 将当前元素添加到 tmp 切片中
		}
	}
	return tmp // 返回存储匹配元素的切片 TMP
}
func DomeFun04() {
	s1 := []string{"nihao", "text", "zhang", "nihao"}
	fmt.Println(strSlice(s1, "nihao"))
}

//5 . 创建一个函数，接受一个整数切片作为参数，返回一个新切片，其中的元素是原切片中相邻两个元素的和。

func FindSlice(s []int) []int {
	tmp := make([]int, 0)
	for i := 0; i < len(s)-1; i++ { // 注意 循环中的条件 i <= len(s) 导致最后一个元素与不存在的下一个元素相加，造成了越界错误
		// 循环条件 i < len(s)-1，确保不会遍历到最后一个元素。这样就能避免越界错误
		tmp = append(tmp, s[i]+s[i+1])
	}

	return tmp
}
func DomeFun05() {
	s := []int{1, 2, 3, 4, 5, 6, 8, 3, 8, 4}
	fmt.Println(FindSlice(s))
}
