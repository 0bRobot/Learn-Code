package days

import (
	"fmt"
)

func Exp13() {
	s := "leetcode"
	n := "leeto"
	fmt.Println(strStr(s, n))

}

func strStr(haystack string, needle string) int {

	k := len(haystack)
	n := len(needle)

	for i := 0; i <= k-n; i++ { // 循环遍历可能的子串
		m := true
		for j := 0; j < n; j++ {
			if haystack[i+j] != needle[j] { // 是在遍历当前子串的每个字符，比较该字符是否与 needle 对应位置的字符不匹配
				m = false // 子串中的某个字符与 needle 不匹配
				break
			}
		}
		if m { // 如果 match 为真，说明当前子串是 needle 的匹配
			return i
		}
	}
	return -1 // 没有找到匹配，返回 -1
}
