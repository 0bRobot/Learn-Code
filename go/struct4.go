package main

import "fmt"

type MyInt int64 // 自定义类型   类型名  新的类型

type MyInt2 = int // 类型别名  使用= 号

// 拓展
type Logger struct {
	Dest string
}

func (back1 Logger) PrintLog() { //  旧方法
	fmt.Println("----测试=====", back1.Dest)
}

type MyLogger Logger //自定义类型

func (back2 MyLogger) PrintVLog() { // 新方法
	fmt.Printf("----\t测试\t=====\n", back2.Dest)
}

func main() {
	// 1.自定义类型  type 关键字
	var NU MyInt = 123 // 长语法
	var (
		NU2 MyInt  = 13
		MU3 MyInt2 = 1212
	)
	fmt.Println(NU, NU2, MU3)
	fmt.Printf("%T,%T,%T", NU, NU2, MU3)
	// 2.类型别名

	// 3 .拓展
	// 旧方法的调用
	p1 := Logger{}
	p1.Dest = "睡吧"
	p1.PrintLog()
	fmt.Println("------------------后--------------------")

	p2 := MyLogger{}
	p2.Dest = "你好"
	p2.PrintVLog()

}
