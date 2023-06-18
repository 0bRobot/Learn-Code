package main

import "fmt"

func main() {
	// 字符和字符串
	y1 := '中' // 字符     20013 // utf-8 编码
	y2 := "中" // 字符串   中     字符串 由字符组成

	fmt.Println(y1, y2)
	fmt.Println("--------------------------")

	//	byte 和 rune 类型   遍历字符串

	fs := "hello World 中国"

	for i := 0; i < len(fs); i++ { // byte    for 按照单个字节遍历
		fmt.Printf("%v(%c)", fs[i], fs[i]) // %T  输出变量的类型

	}
	fmt.Println()

	for _, r := range fs { //rune  for rang 按照单个字符遍历
		fmt.Printf("%v(%c) ", r, r) //  %c  输出单个字符   %u  ；输出无符号十进制数  %v 以默认的方式打印变量的值  自动识别类型
	}

	fmt.Println()

	k1 := "你好呀:小宝贝:666"
	k := 0
	for _, r := range k1 {
		if r == ':' {
			fmt.Println(k)
			break
		}
		k++
	}

}
