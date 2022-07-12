package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//字符与字符串
	fmt.Println("字符与字符串")
	s := "你好,world"
	// 检测变量s 的字符串长度

	// len()  go 语言内置函数 len() 可以求字符串长度

	fmt.Println(len(s)) // 返回长度为 13
	// 字符串底层使用字节来保存
	//len函数 返回字符串字节数

	// byte  字符
	var c byte = 'h' // byte 是无符号的 int8 型   // 打印字符 用单引号
	// var D byte = '你'  // 提示内存溢出
	// go 语言中 非拉丁文如汉语 韩语等使用utf8编码  一个汉字 占3个字节  一个byte 占一个字节
	fmt.Println(c) //输出 104
	fmt.Printf("c= %c\n", c)

	// rune unicode  字符
	var D rune = '看' // rune   int 32  使用4个字节报错
	fmt.Printf("D= %c\n", D)

	//统计汉字字符数
	// 使用  utf8.RuneCountInString() go语言内置函数
	ttt := "特殊123"
	fmt.Println("字符串长度", utf8.RuneCountInString(ttt))

	//  for range 语法
	// Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
	//    for key, value := range oldMap {
	//    newMap[key] = value     								key 和 value 是可以省略
	//     }

	//只想读取 key
	//for key := range oldMap

	for _, r := range s { // 需要两个返回值  _ 占位  s  某个变量
		fmt.Printf("rune类型的变量 %T,%U,%c,\n", r, r, r) // %T 打印类型  %U  打印utf-8编码   %C  打印字符
	}

}
