package _interface

import "fmt"

type Mover interface {
	Move()
}
type Dog struct {
}

// Move 值接受者实现Move方法
//func (d Dog) Move() {
//
//}

// Move 指针接受者实现Move方法
func (d *Dog) Move() {
	fmt.Println("它可以跑")
}

// 使用指针接收者实现接口
// 接口变量可以接受结构体指针类型  但是不能接收值类型 (不是任何值都能取地址 如空值，常量，字面量)

// 使用值接收者实现接口
// 接口变量即能接收指针类型 又能接收值类型 (有了地址就能取值)

// DemoInterface 指针接受者和值接收者
func DemoInterface() {
	var x Mover // 定义接口变量
	d := &Dog{}
	x = d
	fmt.Println(x)

	//n := MyInt(10)
	// n.Move() 	//n 是变量可以取地址

	//MyInt(10).Move()  10 是字面量不能取地址
}

// 案例
type makeDream interface {
	dream()
}
type student struct { // 定义一个结构体 student，表示学生
	name string
	age  int
}

func (s student) dream() { // 为 student 结构体实现 dream() 方法
	fmt.Println("早日成才！！", s.name)
}

type MyInt int // 定义一个新类型 MyInt，基于 int

func (m MyInt) dream() {
	fmt.Println("下辈子做个人！！", m)
}
func DemoInterface1() {
	var s = student{"张三", 32} // 创建一个 student 类型的变量 s
	var x makeDream
	x = s            // 声明一个 makeDream 接口类型的变量 x，将 s 赋值给 x
	x.dream()        // 调用接口的 dream() 方法
	var i MyInt = 10 // 将 i 赋值给 x，此时 x 变为 MyInt 类型
	x = i
	x.dream() // 调用接口的 dream() 方法

	//尽管两个方法的方法名相同，但由于它们分别附加到不同的类型上，因此不会发生冲突
}

//接口组合

//  空接口

type Any interface {
}

// DemoInterface2 空接口
func DemoInterface2() {
	//var x Any
	var x interface{} // 空接口一般都这样定义
	x = 100
	x = "test"
	x = struct{}{}
	fmt.Println(x) // 接口得默认值nil

	//map[string]int  // 	{"name":"李逵","age":19}  与map 对应得数据类型不匹配
	// 空接口类型作为map得value  实现数据据的兼容
	s := map[string]interface{}{
		"name": "张三",
		"age":  39,
	}
	fmt.Println(s) //map[age:39 name:张三]
}

// 接口组合
