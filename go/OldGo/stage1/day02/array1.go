package main

import (
	"fmt"
)

func main() {
	//dArray()
	//d1Array()
	d2Array()
}

func d2Array() {

	a := [3]int{10, 20, 30}

	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [2][2]int{{1, 1}, {1, 1}}

	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}
func modifyArray(x [3]int) { //值传递  把数组a的值赋值给数组x
	x[0] = 100
	fmt.Println(x) // [100 20 30]
}

func modifyArray2(x [2][2]int) {
	x[1][0] = 100
	fmt.Println(x) // [[1 1] [100 1]]
}
