package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//strDome5()
	//strDome6()
	//strDome7()
	//strDome8()
	//strDome9()
	//strDome10()
	strDome11()
}

//结构体内存布局
type strP1 struct {
	a, b, c, d, e int
}

func strDome5() {
	p6 := strP1{
		1, 2, 3, 4, 5,
	}
	fmt.Printf("%p\n", &p6.a) // 0xc00000a300     // &  取内存地址
	fmt.Printf("%p\n", &p6.b) // 0xc00000a308
	fmt.Printf("%p\n", &p6.c) //  0xc00000a310
	fmt.Printf("%p\n", &p6.d) //  0xc00000a318
	fmt.Printf("%p\n", &p6.e) //  0xc00000a320    // %p  打印内存地址
	fmt.Printf("%p\n", &p6)   //  0xc00000a300    结构体地址地址 为第一个字段地址

	fmt.Println(unsafe.Sizeof(p6)) // 40   占用内存大小
}

func strDome6() {
	var anonStr struct{}
	fmt.Printf("%p\n", &anonStr)        // 0x9f2438
	fmt.Println(unsafe.Sizeof(anonStr)) // 0  空结构体是不占用空间的
}

// 构造函数

type pAddr struct {
	name string
	city string
	age  int
}

func newP(name, city string, age int) *pAddr {
	//  声明一个函数  接收2个字符串 一个int类  返回 结构体  * 取值
	return &pAddr{
		//  &pAddr  取函数地址
		name: name,
		city: city,
		age:  age,
	}
	//构造函数返回的是结构体指针类型

}
func strDome7() {

	p10 := newP("小女子", "北京", 23)
	fmt.Printf("%v\n", p10)  // &{小女子 北京 23}
	fmt.Printf("%#v\n", p10) //  &main.pAddr{name:"小女子", city:"北京", age:23}
}

// 方法和接收者
func strDome8() {
	p2 := newP("MrK", "BJ", 23)
	p2.Dream()
}
func (p1 pAddr) Dream() {
	fmt.Printf("你在做梦%v\n", p1.name)
}

//  方法接收者案例

type Person struct {
	//定义一个结构体
	Name string
}

//因为是接收者是值类型，所以该方法不会修改原结构体，而是会返回一个新的Copy

func (p Person) ShowName() Person {
	//定义一个方法，将该方法与结构体Person关联起来，此时Person就是方法的接收者
	return p
	//这里的p是指结构体实例对象

}
func strDome9() {
	//创建结构体实例
	p := Person{Name: "Tom"}
	//调用方法ShowName
	result := p.ShowName()
	//输出结果，注意这里输出的是结构体Person的指针地址
	fmt.Println(&result) //&{Tom}
	//调用后，原实例并未改变
	fmt.Println(p) //{Tom}
}

//  指针类型的接收者

func (p2 *pAddr) setAge(newAge int) {
	p2.age = newAge
}
func strDome10() {
	p1 := newP("张三", "北京", 21)
	fmt.Println(p1.age) // 21
	p1.setAge(40)
	fmt.Println(p1.age) //40
}

//值类型的接收者
func (p pAddr) setAge2(newAge2 int) {
	p.age = newAge2

}
func strDome11() {
	p1 := pAddr{"", "", 30}
	fmt.Println(p1.age)
	p1.setAge2(30)
	fmt.Println(p1.age)
}
