package main

import (
	"fmt"
)

func main() {
	traversalString()

}

func traversalString() {
	// 定义字符串    包含 英文  汉子
	str1 := "|中国红 china 666!!~~"

	//  获取 文字数量个数
	n := 0 // 汉字数量
	m := 0 //  判断字符数量
	l := 0 //  判断数字

	for _, r := range str1 {
		//if unicode.Is(unicode.Han, r) {  判断一个字符 是否是汉字
		//	n++
		//}
		//unicode.IsNumber()  判断一个字符 是否是数字

		//  ascii  码 大写字母 十进制数   65 - 90   小写字母  97 - 122
		//数字    0 - 9   十进制数 为  48 -57

		if r > 256 { // 通过 ASCII 码 的大小值
			//  判断汉字的数量
			n++
			continue

		}
		if r <= '9' && r >= '0' { // 判断数字
			l++
		} else { // 判断字符数量
			m++
		}

	}
	fmt.Println("汉字的数量", n)
	fmt.Println("英文的数量", m)
	fmt.Println("数字的数量", l)

}
