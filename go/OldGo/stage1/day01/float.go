package main

import (
	"fmt"
	"math"
)

func main() {
	f1 := math.MaxFloat32 // float32 最大的浮点数
	f2 := math.MaxFloat64 // float64 最大的浮点数
	f3 := math.Pi         // π的值

	fmt.Printf("float32最大值%f\n", f1) // 打印浮点数时使用  %f  输出  \n 换行
	fmt.Printf("float64最大值%f\n", f2)
	fmt.Printf("π的值%f\n", f3)
	fmt.Printf("π的值%.10f\n", f3) //  %.f  保留几位小数
}
