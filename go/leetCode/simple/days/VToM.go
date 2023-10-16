package days

import "fmt"

func romanToInt(s string) int { // 定义一个函数 romanToInt，接受一个罗马数字的字符串 s，返回对应的整数值

	a := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	r := 0 // 初始化整数结果变量 r 为 0
	p := 0 // 初始化前一个字符的整数值变量 p 为 0

	for i := len(s) - 1; i >= 0; i-- { // 从字符串末尾开始遍历每个字符
		// 核心思想是从右向左遍历罗马数字字符串，根据字符的大小关系进行加法或减法操作，最终得到整数结果
		c := a[s[i]] // 比较当前字符与前一个字符的整数值  s[i] 表示字符串 s 中的第 i 个字符
		if c < p {   // 如果当前字符小于前一个字符，表示减法操作，将 c 从结果中减去
			r -= c
		} else { // 否则，当前字符大于等于前一个字符，表示加法操作，将 c 加到结果中
			r += c
		}

		p = c // 更新前一个字符的整数值为当前字符的整数值，以便下一次比较
	}
	return r // 循环结束后，返回最终的整数结果 r，即为罗马数字字符串 s 所表示的整数值
}

func Exp07() {
	fmt.Println(romanToInt("III"))

}