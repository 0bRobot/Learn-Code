package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//strDome21()
	strDome22()
}

//  结构体

type students struct { //  在函数内部 用type  定义一个结构体
	id       int
	name     string
	age      int8
	mapScope map[string]int //  其中一个字段是  map类型
}

func strDome21() {
	//
	//type students struct { //  在函数内部 用type  定义一个结构体
	//	id       int
	//	name     string
	//	age      int8
	//	mapScope map[string]int //  其中一个字段是  map类型
	//}

	var stu1 students
	stu1.id = 1
	stu1.name = "李莉"
	//stu1.mapScope["语文"] = 100  这里 mapScope map 必须初始化才能使用  否则panic  提示 空map
	stu1.mapScope = make(map[string]int, 5) // 初始化 结构体当中的map  mapScope 字段
	stu1.mapScope["数学"] = 90
	stu1.mapScope["语文"] = 87
	fmt.Printf("%+v\n", stu1) // {id:1 name:李莉 age:0 mapScope:map[数学:90 语文:87]}
	fmt.Printf("%#v\n", stu1) // main.students{id:1, name:"李莉", age:0, mapScope:map[string]int{"数学":90, "语文 ":87}}
	fmt.Println(stu1.name)

	//指针类型结构体
	var user1 students // 声明一个变量 user1 类型是结构体students
	fmt.Printf("%#v\n", user1)
	var user2 = &user1         // user2  存储的是指针， 结构体指针 *students  指向变量user1 的地址内存
	fmt.Printf("%#v\n", user2) // &main.students{id:0, name:"", age:0, mapScope:map[string]int(nil)}

	p1 := new(int)
	fmt.Printf("%v %T\n", p1, p1) //0xc000016228 *int
	p2 := new(students)           // new 返回对应类型的指针
	fmt.Printf("%p %T\n", p2, p2) //&{0  0 map[]}   *main.students
	// %p 打印出指针类型的内存地址   0xc0000726c0 *main.students

}

func strDome22() {
	stu1 := students{}        // 初始化  空
	fmt.Printf("%#v\n", stu1) // main.students{id:0, name:"", age:0, mapScope:map[string]int(nil)}
	stu2 := &students{}       // 取地址  -->  new(students) 返回结构体指针
	(*stu2).age = 17          // 根据地址结构体取值  在.访问 结构体成员
	stu2.name = "张三"        // go语言  语法糖  自动加了一个 *stu2  支持 结构体指针类型.属性  简写
	fmt.Printf("%#v\n", stu2) // 	&main.students{id:0, name:"", age:0, mapScope:map[string]int(nil)}

	// 空结构体
	type dome1 struct{} // 定义一个空结构体

	v1 := dome1{}
	fmt.Println(unsafe.Sizeof(v1)) // 0 空结构体不占用内存空间

	// 空结构体的使用
	set := map[string]dome1{}           //  应用了上面 dome1结构体 初始化为空  map 的 value  为空
	set1 := map[string]struct{}{}       //  使用到了匿名结构体 字段为空 不占用内存空间  map 的 value  为空
	fmt.Printf("%v   %T\n", set, set)   // map[]   map[string]main.dome1
	fmt.Printf("%v   %T\n", set1, set1) // map[]   map[string]struct {}

	// 使用场景  map key占用空间 value  不占空间
	// 如利用map 的key 不能重复的特性 去重

	nameList := []string{"张三", "王五", "赵六", "李四", "王五", "张三", "王莽"} // 定义一个 nameList  切片
	var nameMap = make(map[string]struct{})
	for _, name := range nameList {
		nameMap[name] = struct{}{} // 空结构类型 map 的为空
	}
	for key := range nameMap {
		fmt.Println(key)
	}

}
