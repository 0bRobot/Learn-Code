package main

import (
	"fmt"
)

func main() {
	//strDome12()
	//strDome13()
	//strDome14()
	//strDome15()
	//strDome16()
	strDome17()
}

// 添加一个自定义类型

type myint int

func (m myint) mypPrint() {
	// 为自定义类型添加 mypPrint 方法
	fmt.Println("自定义类型") // 自定义类型
}
func strDome12() {
	var m1 myint
	m1.mypPrint()
	m1 = 100
	fmt.Printf("%#v,%T\n", m1, m1) // 100,main.Myint
}

// 结构体匿名字段

type p01 struct {
	string //只有类型  没有字段名称
	int
}

func strDome13() {
	x := p01{
		"张三",
		23,
	}
	fmt.Printf("%#v\n", x)       // main.p01{string:"张三", int:23}
	fmt.Println(x.string, x.int) //  张三 23
	x.string = "李四"
	fmt.Println(x.string, x.int) // 李四 23
	fmt.Printf("%#v\n", x)       // main.p01{string:"李四", int:23}
}

// 结构体嵌套

type address struct {
	// 地址 结构体
	province string
	city     string
}
type users struct {
	// 用户信息
	id   int
	name string
	Add  address // 类型是结构体 嵌套
}

func strDome14() {
	user1 := users{
		id:   1,
		name: "小明子",
		Add: address{ //  结构体 嵌套
			province: "北京",
			city:     "北京市",
		},
	}
	fmt.Printf("%#v\n", user1) // main.users{id:1, name:"小明子", Add:main.address{province:"北京", city:"北京市"} }
}

// 匿名结构体 嵌套

type addre struct {
	string
	//province string
	city string
}
type users1 struct {
	name   string
	gender string
	addre  // 匿名字段
}

func strDome15() {
	var noUser1 users1
	noUser1.name = "张良"
	noUser1.gender = "男"
	//noUser1.addre.province = "河北省"
	noUser1.string = "河北省" // 匿名字段默认使用类型名作为字段名
	noUser1.city = "石家庄市"  // 匿名字段可以省略

	fmt.Printf("%#v\n", noUser1) // main.users1{name:"张良", gender:"男", addre:main.addre{province:"河北省", city:" 石家庄市"}}
	//main.users1{name:"张良", gender:"男", addre:main.addre{string:"河北省", city:"石家庄市"}}
}

//嵌套结构体的字段名冲突

type Add struct {
	// 地址
	province   string
	city       string
	CreateTime string
}
type email struct {
	// 邮箱
	Email      string
	CreateTime string
}
type use01 struct {
	// 用户
	Name string
	age  int
	Add
	email
}

func strDome16() {
	var U01 use01
	U01.Name = "瞌睡了"
	U01.age = 50
	U01.Add.province = "江苏" // 注意层级关系  测试不影响结构
	//U01.province = "杭州"  // 注意层级关系  测试不影响结构
	U01.Add.city = "南京"
	U01.Add.CreateTime = "2023-03-10" // 指定对应嵌套的结构体
	U01.email.Email = "admin@admin.com"
	U01.email.CreateTime = "2023-03-11" // 指定对应嵌套的结构体
	fmt.Printf("%#v\n", U01)
	//main.use01{Name:"瞌睡了", age:50, Add:main.Add{province:"江苏", city:"南京", CreateTime:"2023-03-10"}, email:main.email{Email:"admin@admin.com", CreateTime:"2023-03-11"}}
	fmt.Println(U01)
	// {瞌睡了 50 {江苏 南京 2023-03-10} {admin@admin.com 2023-03-11}}
}

// 结构体的继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("能跑--%s\n", a.name)
}

type dog struct {
	feet int8
	*Animal
}

func (d *dog) wang() {
	fmt.Printf("会叫 --%s\n", d.name)
}

func strDome17() {
	d1 := &dog{
		feet: 3,
		Animal: &Animal{ // 嵌套的是结构体指针
			name: "旺财",
		},
	}
	d1.wang() // 会叫 --旺财
	d1.move() // 能跑--旺财
}
