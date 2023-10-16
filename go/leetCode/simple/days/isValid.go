package days

import "fmt"

func Exp09() {
	s := "()[]{}"
	fmt.Println(isValid(s))

}

// (  )  [  ]  {    }
// 40 41 91 93 123 125

func isValid(s string) bool {
	ss := []rune{} // 创建一个用于存放括号的栈
	p := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			ss = append(ss, char) // 开括号，推入栈
		} else if char == ')' || char == ']' || char == '}' {
			if len(ss) == 0 {
				return false // 如果没有开括号与之匹配，返回 false
			}
			top := ss[len(ss)-1] // 弹出栈顶元素
			ss = ss[:len(ss)-1]
			e := p[top] // 获取对应的预期闭括号
			if e != char {
				return false //括号不匹配，返回 false
			}
		}
	}
	return len(ss) == 0 // 确保所有括号都正确闭合
}

func isValid1(s string) bool {
	l := len(s) - 1

	for i := 0; i < l; i++ {
		for n := 1; n < len(s); n++ {
			if s[i] == 40 && s[i+n] == 41 {
				return true
			} else if s[i] == 91 && s[i+n] == 93 {
				return true
			} else if s[i] == 123 && s[i+n] == 125 {
				return true
			}
		}
	}
	return false
	// 未处理括号嵌套问题
}
