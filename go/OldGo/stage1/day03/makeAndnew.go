package main

import "fmt"

func main() {
	//mnDome()
	newDome()
	makeDome()
}

func mnDome() {
	var a *int
	//声明一个int类型指针  未分配内存为  空地址 nil
	a = new(int) //  使用 new  分配内存给  int类型的指针    否则panic
	*a = 100     // 所以不能 给改指针赋值
	fmt.Println(a)

	var b map[string]int
	//// map 初始化才能使用
	b = make(map[string]int) //  map  必须使用 make 初始化
	b["1234543"] = 100
	fmt.Println(b)

}

func newDome() {
	//new函数不太常用
	a := new(int)
	b := new(string)
	//new函数返回的是一个该类型的指针
	fmt.Printf("%T %v\n", a, a) // *int 0xc0000aa058
	fmt.Printf("%T %v\n", b, b) // *string 0xc000088220
	// 该指针对应的值为该类型的零值
	fmt.Printf("%v\n", *a) //0
	fmt.Printf("%v\n", *a) //0

}

func makeDome() {
	//只用于slice、map以及channel的内存创建
	//返回的类型就是这三个类型本身
	b := make(map[string]int, 20)
	b["today"] = 2023
	fmt.Println(b)

	c := make([]string, 0, 8)
	//使用make  初始化slice
	d := append(c, "测试1", "第二")
	fmt.Println(d)
}
