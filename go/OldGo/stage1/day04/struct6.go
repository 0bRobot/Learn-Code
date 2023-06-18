package main

import (
	"fmt"
	"unsafe"
)

func main() {
	strDome23()
}

// 结构体内存大小

type MemoryDome1 struct {
	A int8 // 占用一个字节  1 byte
	B int8
	C int8
}

type MemoryDome2 struct {
	a int8  //  1 byte
	b int32 //  4 byte
	c int8  //  1 byte
	d int64 //  8 byte
}
type MemoryDome3 struct {
	a int8  //  1 byte
	c int8  //  1 byte
	b int32 //  4 byte
	d int64 //  8 byte
}

type MemoryDome4 struct {
	aa int8     // 1 byte
	bb struct{} // 0
	// 当空结构体类型作为结构体的最后一个字段时，如果有指向该字段的指针，那么就会返回该结构体之外的地址。为了避免内存泄露会额外进行一次内存对齐。
	// 所以不建议一个空结构体 放到结构体的最后字段
}

func strDome23() {
	// 结构体是占用一块连续的内存，一个结构体变量的大小由结构体中的字段决定
	var ff MemoryDome1
	fmt.Println(unsafe.Sizeof(ff)) //3  几个字段的累加

	var ss MemoryDome2
	fmt.Println(unsafe.Sizeof(ss)) //  24 byte  大小为什么是24字节   不是累加 ?
	//  结构体的内存对齐策略    结构体的大小 又不完全由结构体的字段决定
	// 结构体按照结构体字段最大字节 进行内存对齐  对齐系数
	// 在了解了Go的内存对齐规则之后，我们在日常的编码过程中，完全可以通过合理地调整结构体的字段顺序，从而优化结构体的大小。

	var ss1 MemoryDome3
	// 调整结构体字段后  利用cpu读取的粒度  优化结构体大小及 减少cpu 读取次数
	fmt.Println(unsafe.Sizeof(ss1)) // 16 byte

	var s2 MemoryDome4
	fmt.Println(unsafe.Sizeof(s2)) // 2  ?
}
