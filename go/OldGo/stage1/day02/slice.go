package main

import "fmt"

func main() {
	//dSlice()
	d1Slice()
	//d2Slice()
}

func dSlice() {
	// slice 的声明

	var x []int
	var x1 []string
	var x2 []bool
	var x3 = []int{}
	var x4 = []int{}
	fmt.Printf("%#v \n", x)  // []int(nil)
	fmt.Printf("%#v \n", x1) //  []string(nil)
	fmt.Printf("%#v \n", x2) //  []bool(nil)
	fmt.Printf("%#v \n", x3) //  []int{}
	//fmt.Println(x4 == x3)  切片是不能比较的
	fmt.Println(len(x4)) //求切片的长度(元素个数)
	fmt.Println(cap(x4)) //求切片的容量(底层数组能存的元素个数)

}

func d1Slice() {

	A := [...]int{1, 2, 3, 4, 5, 6, 7} //声明一个数组
	s := A[1:3]                        // 指定low和high两个索引界限值声明并初始化切片    索引从1切到3   左包含，右不包含
	fmt.Printf("%#v \n", s)            // []int{2, 3}
	fmt.Println(len(s), cap(s))        // 切片s 长度2  容量 6
	s1 := A[:5]                        // 索引从0 开始切到索引5  不包含索引5    左包含，右不包含
	fmt.Println(len(s1))               //5
	fmt.Println(cap(s1))               //7
	s2 := A[4:]                        // 索引从4 开始切到索引最后
	fmt.Println(len(s2))               //3
	fmt.Println(cap(s2))               //3
	s3 := A[:]                         //索引从0 开始切到索引最后
	fmt.Println(len(s3))               //7
	fmt.Println(cap(s3))               //7

	str1 := "Hello world"
	B := str1[:4]
	fmt.Printf("%T \n", B) // 类型是 string
	fmt.Println(B, len(B)) // Hell  长度是 4
	//fmt.Println(B, cap(B))      // 字符串切片没有容量这个概念

	AAA := []int{0, 12, 3, 4, 6}
	fmt.Println(AAA)
	BBB := []int{3: 10, 5: 200}
	fmt.Println(BBB)

}

func d2Slice() {
	//切片在切片

	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//         0  1  2  3  4  5  6   7  8
	B := A[1:4]                 //      左包含右不包含
	fmt.Println(len(A), cap(A)) //9 9

	fmt.Println(B)              //[2 3 4]
	fmt.Println(len(B), cap(B)) // 3 8

	C := B[:5]                  //切片在切片，high的取值范围<= cap
	fmt.Println(C)              //[2 3 4 5 6]
	fmt.Println(len(C), cap(C)) //5 8

	D := C[3:5]
	fmt.Println(D)              //[5 6]
	fmt.Println(len(D), cap(D)) //2 5
	//切片在切片时  修改新切片的值 原始切片的值也会发生改变
	//底层操作的同一个数组的数据
	D[1] = 200
	fmt.Println(D) //[5 200]
	fmt.Println(A) // [1 2 3 4 5 200 7 8 9]
	fmt.Println(C) //[2 3 4 5 200]

}
