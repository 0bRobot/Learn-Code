package main

import "fmt"

func main() {
	strDome24()
}

// 方法
// Go语言中的`方法（Method）`是一种作用于特定类型变量的函数。这种特定类型变量叫做`接收者（Receiver）`

type Person01 struct {
	name string
	age  int
}

// dream 给 person01 定义一个方法
//2 方法 dream  只有 person01 类型的人 才能调用该方法
func (p Person01) dream(s string) { // 第一个()  给那个类型添加方法 类型的首字母小写作为接收者
	// 第二个() 传入的参数 及传入的类型
	fmt.Printf("%s的梦想是%s\n", p.name, s)
}

//func (p Person01) guonian3() {
//	p.age = p.age + 3  // 只修改了方法中值拷贝的副本 原结构体中的值并为修改   所以要要用指针类型的接收者 修改结构体中变量的值
//}
func (p *Person01) guonian3() {
	p.age = p.age + 3
}

// 结合构造函数
func newPerson(name string, age int) Person01 {
	return Person01{
		name: name,
		age:  age,
	}
}

func strDome24() {
	//方法 dream  只有 person01 类型的人 才能调用该方法

	// 1 结构构造函数 调用该方法
	x1 := newPerson("张三", 345)
	x1.dream("玩玩")

	// 2.不用构造函数
	x2 := Person01{
		name: "王五",
		age:  22,
	}
	x2.dream("打游戏")
	x2.guonian3()       // 调用guonaian3方法 年龄+3
	fmt.Println(x2.age) // 用指针类型的接收者 修改结构体中变量的值
}
