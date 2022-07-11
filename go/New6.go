package main

import "fmt"

func main() {
	//1.6 循环语句

	// ++ 自增操作符  将变量的值自增  +1
	// -- 自减操作符   将变量的值自减 -1
	A := 5
	A++
	B := 8
	B--
	fmt.Println(A)
	fmt.Println(B)
	// 循环1
	// 循环的初值    循环结束的条件  循环增量或减量
	for i := 1; i <= 10; i++ {
		//循环体
		fmt.Println("输出1-10： ", i)

	}
	// 循环2   循环的增量或减量 咋循环体内
	for O := 1; O <= 10; {
		fmt.Println(O)
		O++ //循环的增量或减量 咋循环体内
	}
	// 循环3 死循环
	//for {  //恒为真   死循环
	//	fmt.Println(123)  // 无限循环 死循环
	//}

	// 打印 1 - 9 但是跳过 5
	for c := 1; c < 10; c++ {
		if c == 5 {
			// break  退出循环
			continue //跳过当前循环
		}
		fmt.Println("continue AND  break ", c)
	}
}
