package days

import "fmt"

func longestCommonPrefix(strs []string) string {
	r := ""             // 初始化结果字符串
	if len(strs) == 0 { // 如果输入切片为空，直接返回空字符串
		return ""
	}

	min := len(strs[0]) // 初始化最小字符串长度为第一个字符串的长度

	for _, v := range strs { // 遍历输入切片，找到其中最短字符串的长度
		if len(v) < min {
			min = len(v)
		}
	}

	for i := 0; i < min; i++ { // 遍历每个字符的位置
		char := strs[0][i] // 取出第一个字符串的当前位置的字符

		for _, s := range strs { // 检查其他字符串的相同位置是否都为相同字符
			if s[i] != char {
				return r // 如果有不同，返回当前结果字符串
			}
		}
		r += string(char) // 如果所有字符串在当前位置都相同，将字符追加到结果字符串
	}
	return r
}

func Exp08() {
	s := []string{"a"}

	fmt.Println(longestCommonPrefix(s))
}
