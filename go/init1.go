package main

import (
	"fmt"
	"net"
)

func main(){

// 变量的声明
	// 1.标准格式
	// var 关键字   变量名   变量类型   值   四个部分
	var a int = 3  // 声明变量a
	var c,d int = 2,3
	var e,f = 3,4   // 自动推导类型
	var r,g =2,3
	// 类型
	// bool
	//string
	//int[一般是4个 随系统]   int8[一个字节]  int16[两个字节]  in32[4个字节]  int64[8个字节]    一个字节  8位  也就是 8bit
	// uint(无符号整型)   uint8[一个字节]  uint16[两个字节]  uin32[4个字节]  uint64[8个字节]   uintptr
	var wrwr uintptr
	// float32    float64
	// complex64  complex128

	// byte // uint8 的别名      uint8 范围 1-255  int8 范围  -128-- 127
	// rune // int32 的别名  代表一个Unicode码

	// 当一个变量边声明后 系统自动赋值该类型的零值
	// int 为0    float 为 0.0  bool  为false  string为空字符串   指针为nil

	// 所有的内存在go中都是经过初始化的
	fmt.Println(a,c,d,e,f,r,g,wrwr)

	fmt.Printf("%d\n",a)  //
	fmt.Printf("%T\n",wrwr)   // %T  输出类型  %d  整数占位付

// 批量声明变量

	var (
		aa int
		bb = 3.14
		cC = "你好啊"
	)
	fmt.Printf("aa=%d bb=%f c=%s\n",aa,bb,cC)


// 简短格式 声明变量

	// --------缺点
	// 1. 变量定义 同时显示初始化
	// 2. 不能提供数据类型
	// 3. 只能用在函数内部
	TT:= 324
	GG:="Hello MrYang"
	fmt.Println(TT,GG)
	fmt.Printf("%T,\r1\n",GG)

// 多重赋值
	var  level = 23
	// level :=23  提示 :=左侧没有新变量  //  多重赋值
	fmt.Println(level)

	var conn net.Conn
	var err error

	conn,err = net.Dial("tcp","127.0.0.1:48450")  // net.Dial  有两个返回值
	fmt.Println(conn,err)

	conn1,err := net.Dial("tcp","127.0.5.1:38450")   // err 接收错误   相当于 err重复赋值

	fmt.Println(conn1,err)
	// 结论  -------------------------------   重复定义的例外
	// 多个短变量声明个赋值中，至少有一个新声明的变量出现在左值中，即便其他变量名 可能是重复声明的，编译器也不会报错

fmt.Println("--------------------------Demo----------------------------------------------------")

// 变量交换 Demo

		// 实现1
		//var jj = 100
		//var ii = 200
		//var kk int
		//kk = ii
		//ii = jj
		//jj = kk
		//fmt.Println("第一种")
		//fmt.Println(jj,ii)

		fmt.Println("---------------------第二种---------------------------------")
		//var jj = 100
		//var ii = 200
		//jj = jj^ii
		//ii = ii^jj
		//jj = jj^ii
		//fmt.Println("第二种")
		//fmt.Println(jj,ii)


	fmt.Println("---------------------第三种---------------------------------")
	var jj = 100
	var ii = 200
	ii,jj = jj,ii //  go语言独有的变量交换表达式 -----------------------------

	fmt.Println("第三种")
	fmt.Println(jj,ii)

// 匿名变量
		// 使用多重赋值时， 如果不需要在左值中接收变量， 可以使用匿名变量

	//conn4, _ := net.Dial("udp","128.43.23.2:9090")
	conn3,_ := net.Dial("udp","128.43.23.2:9090")
		// 如果不想接收 err的值  可以使用 _表示  这就是 匿名变量
	fmt.Println(conn3)
	// 匿名变量 以_ 表示
	// 匿名变量不占用命名空间，也不会分配内存。
	// 匿名变量可以重复声明使用
	// _ 本身就一个一个特殊的标识符，成为空白标识符。它可以向其他标识符那样用于变量声明或赋值(任何类型都可以赋值给它)
	//但任何赋值给这个标识符都将被抛弃。 不能对它就行  赋值或运算

	//匿名变量单独直接开头
 	// _ := 1

// 作用域
	//一个变量(常量、类型或者函数) 在程序中都有一定的作用范围  称为作用域。
	// go 语言[静态语言]会在编译是检测每个变量是否被使用过，一旦出现未使用的变量，编译器就会报错

	// 变量定义的位置不同
	// 函数内部定义的变量 称为 局部变量    函数的参数和返回值变量都属于局部变量   函数调用结束后这个局部变量就会被销毁
	// 函数外部定义的称为去全局变量   全局变量只需要在一个源文件中定义 就在所有源文件中可以使用，  全局变量 必须 以var 关键字开头   想在外部包中使用 首字母必须大写
	// go语言中 全局变量与局部变量可以相同  但是函数体内的局部变量会被优先考虑

	// 函数定义中的变量称为形式参数 简称形参





}
