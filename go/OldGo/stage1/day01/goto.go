package main

import "fmt"

func main() {

	//break1()
	//goto1()
	break2()
	continue1()
}
func goto1() {
	for i := 0; i < 10; i++ { // 第一层循环

		for j := 0; j < 10; j++ { // 第二层循环
			if j == 2 { // if 判断 条件满足   执行 goto语句
				goto breakLabel // 设置退出标签
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

	// goto 标签
breakLabel: // 标签名 加:号  先定义  Label
	fmt.Println("结束for循环")
}

func break1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true // flag 设置为 true
				break            //  退出 第二层循环
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag { // 判断为true   执行 第二次 break
			break // 退出第一层循环
		}
	}
}

func break2() {
BREAKDEMO1: //  break  标签 只能定义在对应的  for 、switch  、select 代码块上
	// 退出for 循环 按顺序 向下执行
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("break 退出")
}

func continue1() {
Lable:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				continue Lable //  i==2  j==2 条件满足  退出到  Lable  重新循环
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

}
