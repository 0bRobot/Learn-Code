package days

import "fmt"

func Exp21() {
	fmt.Println(isHappy(19))

}

func isHappy(n int) bool {
	s := make(map[int]bool) // 用于存储已经出现过的数字

	for {
		// 如果当前数字已经在集合中出现过，表示陷入循环，不是快乐数
		if s[n] {
			return false
		}
		s[n] = true
		newNumber := 0
		for n > 0 {
			d := n % 10                 //n % 10 操作用于获取 n 的最后一位数字，即取模运算 (n % 10)，结果就是 n 的最后一位数字（个位数）。
			newNumber = newNumber + d*d // newNumber = newNumber + d * d 将每位数字的平方累加到 newNumber 中。
			n = n / 10                  // n = n / 10 将 n 除以10，实际上是将 n 的最后一位数字去掉，因为整数除以10相当于右移一位，去掉了个位数字
		}
		if newNumber == 1 {
			return true
		}
		n = newNumber // 更新当前数字为新数字，继续下一轮循环
	}

}
