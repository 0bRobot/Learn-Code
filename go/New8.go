package main

import "fmt"

func main() {
	// 1.8 指针
	// 指针 是用来保存内存地址的一种特殊变量类型
	//一个指针变量指向了一个值的内存地址。

	// 指针的定义
	var AA *int
	var AA1 *float32
	var C int

	fmt.Println(AA)  //指针默认初始值 是 nil
	fmt.Println(AA1) // nil 不是go语言的关键字    nil是一个种特殊的类型 只能用于  指针  ， 函数  管道  切片  ，interface ，map 的初始值
	fmt.Println(C)   //普通变量的默认初始值是 0   nil 指针也称为空指针
	// nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

	//if(ptr != nil)     /* ptr 不是空指针 */
	//if(ptr == nil)    /* ptr 是空指针 */

	fmt.Printf("指针的类型%T\n", AA) //*int
	fmt.Printf("指针的地址%p\n", AA) // 0x0   16进制  %p 打印指针地址
	// 获取 指定 地址
	//& 取地址操作符
	n := 123
	AA = &n                     // 把 变量n 在内存中的地址 赋值给AA
	fmt.Println(AA)             //     输出变量n 在内存中的地址  0xc0000140e0
	fmt.Printf("变量的地址%p\n", AA) // 只能打印 指针类型

	// * 解指针  操作符
	fmt.Println(*AA) //  输出 变量  n 的值  123
	*AA = 200
	fmt.Println(*AA) //  输出 指针新赋值后的值  200
	fmt.Println(n)   //  变量n 的值 已经被改变

	// 注意
	//----------------------------------
	//var P3 *int      // 指针没有初始化的时候 默认是没有分配内存的   不能进行 取值/解指针操作
	//fmt.Println(*P3)

	// new() 初始化指针
	var p4 *int
	p4 = new(int)
	fmt.Println(*p4)                 //  初始化值为0
	fmt.Printf("new()初始化地址%p\n", p4) //new()初始化地址0xc000014100  已经有了 地址

}
