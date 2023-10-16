package days

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	sum := 0                        // 初始化变量用于记录最后一个单词的长度
	if strings.TrimSpace(s) == "" { // 检查输入字符串是否全为空格
		return 0
	}
	i := len(s) - 1   // 从字符串末尾开始向前遍历
	for s[i] == ' ' { // 跳过末尾的空格
		i--
	}
	for i >= 0 && s[i] != ' ' { // 计算最后一个单词的长度
		sum++
		i--
	}
	return sum
}

func Exp11() {

	s := "  "

	fmt.Println(lengthOfLastWord(s))

}
