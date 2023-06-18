package main

import (
	"fmt"
)

func main() {

	//runDome()
	//runDome1()
	//runDome2()
	//runDome3()
	//initStr()
	//init1Str()
	//init2Str()
	init3Str()
}

// 结构体的定义
type strDome1 struct {
	//  strDome1    自定义结构体的名称，在同一个包内不能重复
	name       string //  name  字段名  后跟类型
	age        int
	addr, mail string // 多个字段同一类型 可以写在一行
}

func runDome() {
	// 结构体的实例化
	var str1 strDome1    // 结构体本身也是一种类型，可以像声明内置类型一样使用var关键字声明结构体类型。
	str1.name = "疯子zhao" //  通过.来访问结构体的字段（成员变量）
	str1.age = 100
	str1.addr = "北京"
	str1.mail = "74108520@gmail.com"
	fmt.Println(str1)         //  {疯子zhao 100 北京 74108520@gmail.com}
	fmt.Printf("%v\n", str1)  // {疯子zhao 100 北京 74108520@gmail.com}
	fmt.Printf("%#v\n", str1) // main.strDome1{name:"疯子zhao", age:100, addr:"北京", mail:"74108520@gmail.com"}
}
func runDome1() {
	// 匿名结构体
	//注意:匿名结构体是var声明变量， 普通结构体是type在函数外部
	var anonSrt struct {
		name, addr string
		age        int
	}
	anonSrt.name = "anonymousName"
	anonSrt.addr = "unknown"
	anonSrt.age = 100
	fmt.Println(anonSrt) // {anonymousName unknown 100}
}

// 创建指针类型结构体
type class1 struct {
	name string
	age  int
}

func runDome2() {
	var p1 = new(class1)
	fmt.Printf("%v\n", p1) // &{ 0}
	// 在Go语言中支持对结构体指针直接使用.来访问结构体的成员
	p1.name = "赵同学"
	p1.age = 12
	fmt.Printf("%v\n", p1)  // &{赵同学 12}
	fmt.Printf("%#v\n", p1) // &main.class1{name:"赵同学", age:12}
	addr1 := &p1.name       // 对函数成员 取内存地址
	fmt.Println(addr1)      // 0xc000004078
	*addr1 = "小明同学"         // 对该地址 进行赋值
	fmt.Println(p1)         //  &{小明同学 12}
}

// 取结构体地址实例化
type addrStr struct {
	name    string
	age, id int
}

func runDome3() {
	p2 := &addrStr{}
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	fmt.Println(p2)             // &{ 0 0}
	fmt.Println(p2.id == p2.id) // 结构体成员类型相同 成员可以进行比较
	fmt.Println(p2.id + p2.age) //   结构成员体类型相同 成员可以进行运算

	p2.name = "小甜甜"
	p2.id = 19
	p2.age = 20
	fmt.Println(p2)              // &{ 0 0}      &{小甜甜 20 19}
	fmt.Println(p2.id == p2.age) // 结构体成员类型相同 成员可以进行比较  false
	fmt.Println(p2.id + p2.age)  //   结构体成员类型相同 成员可以进行运算  39
	// fmt.Println(p2.name == p2.age)     成员类型不同
}

//结构体初始化
type initDemo struct {
	// 定义结构体
	name, city string
	id, age    int
}

func initStr() {
	var p3 initDemo
	//  没有初始化的结构体，其成员变量都是对应其类型的零值
	fmt.Printf("%#v\n", p3) // main.initDemo{name:"", city:"", id:0, age:0}
}

func init1Str() {
	//  使用键值对对结构体进行初始化
	p3 := initDemo{
		name: "小可爱",
		id:   12,
		age:  01,
		// city 字段  未初始化 该类型的零值
		//	当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	}
	fmt.Printf("%#v\n", p3) // main.initDemo{name:"小可爱", city:"", id:12, age:1}
}
func init2Str() {
	//对结构体指针进行键值对初始化
	p4 := &initDemo{
		name: "瞌睡了",
		city: "未知",
		age:  123,
		//id  当某些字段没有初始值的时候，该字段可不写。
	}
	fmt.Printf("%#v\n", p4) // &main.initDemo{name:"瞌睡了", city:"未知", id:0, age:123}
}
func init3Str() {
	//	结构体指针进行值初始化     直接写值
	p5 := &initDemo{
		"该睡觉了",
		"上海",
		01,
		18,
	}
	fmt.Printf("%#v\n", p5) // &main.initDemo{name:"该睡觉了", city:"上海", id:1, age:18}

	// 结构体进行值初始化   直接写值
	p6 := initDemo{
		"马上睡", // 直接写值
		"北京",  // 必须写上所有字段
		02,    //  不能和混用
		21,    // 顺序必须一致
	}
	fmt.Printf("%#v\n", p6) // main.initDemo{name:"马上睡", city:"北京", id:2, age:21}
}
