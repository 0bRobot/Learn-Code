package main

import (
	"fmt"
	"math"
	"unsafe"
)

//在静态类型语言中  (C++/java/Golang等) 中规定在创建一个变量或者常量是，必须指定
//出响应的数据类型，否则无法个变量分配内存

//整形分两大类
// 有符号的整形  int 8,  int 16 , int 32  , int 64,int
// 无符号的整形   uint 8 ， uint 16, uint 32 ,uint 64 ,uint

//Integer 有符号的整形
func Integer() {
	var num8 int8 = math.MaxInt8
	var num16 int16 = math.MaxInt16
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	var num int = math.MaxInt

	fmt.Printf("num8的类型是 %T,num8的大小是%d,mun8是 %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是 %T,num16的大小是%d,mun8是 %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是 %T,num32的大小是%d,mun8是 %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是 %T,num64的大小是%d,mun8是 %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是 %T,num的大小是%d个字节,mun8是 %d\n", num, unsafe.Sizeof(num), num)
}

func unInteger() {
	var unum8 int8 = math.MaxInt8
	var unum16 int16 = math.MaxInt16
	var unum32 int32 = math.MaxInt32
	var unum64 int64 = math.MaxInt64
	var unum int = math.MaxInt

	fmt.Printf("unum8的类型是 %T,unum8的大小是%d,mun8是 %d\n", unum8, unsafe.Sizeof(unum8), unum8)
	fmt.Printf("unum16的类型是 %T,unum16的大小是%d,mun16是 %d\n", unum16, unsafe.Sizeof(unum16), unum16)
	fmt.Printf("unum32的类型是 %T,unum32的大小是%d,mun32是 %d\n", unum32, unsafe.Sizeof(unum32), unum32)
	fmt.Printf("unum64的类型是 %T,unum64的大小是%d,mun64是 %d\n", unum64, unsafe.Sizeof(unum64), unum64)
	fmt.Printf("unum的类型是 %T,unum的大小是%d个字节,mun8是 %d\n", unum, unsafe.Sizeof(unum), unum)
}

// 浮点型
func ShowFloat() {
	//浮点数的精度
	var num1 float32 = math.MaxFloat32
	var num2 float64 = math.MaxFloat64
	fmt.Printf("num1的类型是%T,mun1是%g\n", num1, num1) // float32的精度只能提供大约6个十进制数(表示小数点后6位)的精度
	fmt.Printf("num1的类型是%T,mun1是%g\n", num2, num2) // float64的精度只能提供大约15个十进制数(表示小数点后15位)的精度

}

// 字符型
// 字符串中的每一个元素叫作“字符”，定义字符时使用单引号

//Go 语言的字符有两种
// byte   表示 UTF-8 字符串的单个字节的值，表示的是 ASCII 码表中的一个字符，uint8 的别名类型    1个字节数
// rune   表示单个 unicode 字符，int32 的别名类型                                       4个字节数

func showChat() {
	var xx byte = 65
	var yy uint8 = 65
	fmt.Printf("xx=%c\n", xx)
	fmt.Printf("yy=%c\n", yy)
	var Size_x byte = 65
	fmt.Printf("Size_d,占%d个字节 %c\n", unsafe.Sizeof(Size_x), Size_x)
	var Size_y rune = 'A'
	fmt.Printf("Size_d,占%d个字节 %c\n", unsafe.Sizeof(Size_y), Size_y)
	var HZ rune = '前'
	fmt.Printf("Size_d,占%d个字节 %c\n", unsafe.Sizeof(HZ), HZ)
	//byte 类型只能表示 28个值，所以你想表示其他一些值，例如中文的话，就得使用 rune 类型
}

func main() {
	//一般是用 int表示整形的宽度， 在32位系统下就是32位， 64位系统就是64位
	Integer()
	fmt.Println("================整形==========================")
	unInteger()
	fmt.Println("------------------浮点型--------------------------")
	ShowFloat()
	fmt.Println("--------------------字符----------------------")
	showChat()
}
