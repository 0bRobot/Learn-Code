package main

import "C"
import "fmt"

//import "fmt"
//
//func ttt() { // 函数
//	fmt.Println("测试")
//}
//
//type ( // 定义多个结构
//	no1 struct { //  1
//		name string
//	}
//	no2 struct { // 2
//		id int
//	}
//)
//
//// 结构体方法  接收者
//func (aa no1) test() { //定义 func(接收者) name(返回值)  name 方法名称
//	fmt.Println("测试", aa.name)
//}
//func (bb no2) test1() {
//
//}
//
//func main() {
//	// 结构体的方法
//
//	// 函数 是一种业务逻辑或者功能的封装
//
//	//初始化结构体
//	Dno2 := no2{id: 123}
//	Dno2 := no2{id: 313}
//
//}

//定义函数功能  // 普通函数
//func HS1() {
//	fmt.Println("义")
//}

// 定义两个机构体

type ( //定义两个机构体  AA   BB
	AA struct {
		name string
	}
	BB struct {
		name string
	}
)

//结构体的方法  接收者
// 值接收者
func (C AA) DZ() { // 使用func 关键字    func(接收者 为那个结构体定义的方法) name(返回值)  name 方法名称
	fmt.Printf("传入之时%p\n", &C) // %p 取地址
	C.name = "可乐"
	fmt.Printf("%s 义\n", C.name)
}

func (C BB) DZ() { // 使用func 关键字    func(接收者 为那个结构体定义的方法) name(返回值)  name 方法名称
	fmt.Printf("%s 义\n", C.name)
}

func (E *BB) DZ1() { //  指针接收者  改变值
	fmt.Printf("传入之时%p\n", &E) // %p 取地址
	E.name = "白石"
	fmt.Printf("%s 义\n", E.name)
}

func main() {

	// 函数 是一种特定业务逻辑或者功能的封装
	// 方法 讨论它的对象 对象就是结构体
	fmt.Println("结构体")
	c1 := BB{name: "小飞侠"}
	fmt.Printf("传入之前%p\n", &c1) // %p 取地址
	c1.name = "小飞侠"
	c1.DZ()
	fmt.Printf("%#v", c1)
	c2 := BB{name: "不知道"}
	c2.DZ()
	// 结构体的可见性
	// 变量 函数名称首字母大写为公用的(对于外部其他包是可见的)  不可见 小写 只能在main包下使用
	// 对于的成员名称也需要大写 否则访问不到
}
