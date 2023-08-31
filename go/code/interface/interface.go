package _interface

import "fmt"

type Mover interface {
	Move()
}
type Dog struct {
}

// Move 值接受者实现Move方法
func (d Dog) Move() {

}

// Move 指针接受者实现Move方法
//func (d *Dog) Move() {
//}

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

}

func DemoInterface1() {

}

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
