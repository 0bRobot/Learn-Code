package main

import "fmt"

func main() {
	strDome25()
}

// 结构体的继承

type Animal1 struct {
	name string
}

// 给Animal1 构造一个方法
func (a Animal1) move() {
	fmt.Printf("%s会动...\n", a.name)
}

// 定义一个狗的结构体
type dog1 struct {
	Animal1 // Animal  嵌入结构体来实现类似继承的效果
	leg     int
}

func (d *dog1) wang() {

	fmt.Printf("%s汪汪汪！！\n", d.name)
}

func strDome25() {
	var aa = dog1{
		leg: 4,
		Animal1: Animal1{ // 结构体嵌套
			name: "旺财",
		},
	}
	aa.wang()
	aa.move()
}
