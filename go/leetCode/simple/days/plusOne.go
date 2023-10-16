package days

import "fmt"

func Exp15() {
	s := []int{9, 7}

	fmt.Println(plusOne(s))
}

func plusOne(digits []int) []int {
	n := len(digits)
	// 从最后一位开始加一
	for i := n - 1; i >= 0; i-- {
		// 如果当前位小于9，直接加一返回
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			// 如果当前位是9，变为0，继续下一位
			digits[i] = 0
		}

	}

	// 如果所有位都是9，需要在数组最前面加一个进位1
	digits = append([]int{1}, digits...)
	return digits
}
