package main

import "fmt"

func main() {

	// 类型转换
	// T()
	var i11 int8 = 31
	fmt.Println(int64(i11)) // init 8 转换为 int64  int64 所占空间较大可以转换int8

	var i12 float64 = 142.321 //  float 64 转  int64
	i13 := int64(i12)
	fmt.Println(i13) // 精度 丢失 小数点 消失
	fmt.Printf("%T \n %T", i12, i13)

}
