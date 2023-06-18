package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//pointerDome()
	//pointerDome1()
	//pointerDome2()
	pointerDome3()
}
func pointerDome() {
	// 指针取址
	pv := "hello  pointer"   // pv 是一个字符串变量
	dr := &pv                //  把pv变量的内存地址 存储到变量dr 中
	fmt.Printf("址%v\n", dr)  // 址0xc00003c250
	fmt.Printf("类型%T\n", dr) // 类型*string    类型为 字符串类型的指针

	pv1 := 123 // pv 是一个整形变量
	dr1 := &pv1
	fmt.Printf("址%v\n", dr1)  //  址0xc0000160d0
	fmt.Printf("类型%T\n", dr1) //  类型*int  int 指针

	pv2 := [3]int{1, 2, 3} // pv2 数组
	dr2 := &pv2
	fmt.Printf("址%v\n", dr2)  //  址 &[1 2 3]
	fmt.Printf("类型%T\n", dr2) //   类型 *[3]int   [3]int 类型的指针

}

func pointerDome1() {
	//指针取值
	no1 := 10                   // 定义个int 类型的变量
	a := &no1                   // 取变量no1的地址，将指针保存到a中
	fmt.Printf("变量a类型 %T\n", a) // 变量a类型 *int
	b := *a                     // 指针取值(根据变量a中的内存地址 取值)
	fmt.Printf("b 的类型%T\n", b)  // b 的类型int
	fmt.Printf("b 的值%v\n", b)   // b 的值10
}
func modify(x *int) {
	// modify 函数 接收一个int类型的指针
	// 对指针地址的值 进行操作(重新赋值)
	*x = 200
}
func pointerDome2() {
	a := 10
	modify(&a)            // 取地址 传入函数
	fmt.Printf("%v\n", a) // 200
	fmt.Printf("%T\n", a) // int

}

func pointerDome3() {
	d := [...]int{1, 2, 3}
	p := &d

	p2 := (*int)(unsafe.Pointer(p))

	p2 = (*int)(unsafe.Add(unsafe.Pointer(p2), unsafe.Sizeof(p[0])))
	*p2 += 100
	fmt.Println(d) // [1 102 3]

}
