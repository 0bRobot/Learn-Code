package main

import "fmt"

func main() {
	f31()
}

func f31() {
	//先入后出
	//f31函数结束是运行
	// recover  捕获panic
	defer func() {
		fmt.Println(1235456)
		r := recover() // 通过recover将程序恢复回来，继续往后执行
		fmt.Println("recover----", r)
	}()
	var mm map[string]int
	//panic("手动panic  以下 2行代码 不会执行 ") //  defer  recover 任会处理异常
	mm["张撒"] = 100 //  map 没有初始化 不能直接使用  否则 就出出现 panic
	fmt.Println("panic 之前，这条不会执行")

}
