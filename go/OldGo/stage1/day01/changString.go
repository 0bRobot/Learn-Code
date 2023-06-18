package main

import "fmt"

func main() {
	q1 := "big"
	// 强制类型装换
	byteS1 := []byte(q1) //  把big 转换成字节切片类型
	byteS1[0] = 'p'      //  把第一个切片的字符 修改成 p
	fmt.Println(string(byteS1))
	fmt.Println()

	q2 := "小可爱"
	runeS2 := []rune(q2) //  把小可爱 转换成rune切片类型
	runeS2[0] = '大'      // 把 第一个字 改成大
	fmt.Println(string(runeS2))

}
