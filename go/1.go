package main // package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包
import "fmt"

// import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），
//fmt 包实现了格式化 IO（输入/输出）的函数
func main() { //{ 不能在单独的行上
	// func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，
	// 一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）
	fmt.Println("hello Mr_Robot")
	// fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
	//使用 fmt.Print("hello, world\n") 可以得到相同的结果。
	//Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。
	fmt.Println("你好啊", "小可爱")
	fmt.Println("你好啊" + "小可爱")
	// var Tt = "hello" //声明变量的一般形式是使用 var 关键字
	var aAge, bAge float32 = 32, 55.23
	var A_name string = "Mr_Robot"
	var B_name string = "BBBBB"
	var a, b, c, d int = 2, 6, 7, 85
	var (
		AA, BB int = 2, 3
		CC     int = 5
	)
	fmt.Printf("%s,%f,%s,%f\n", A_name, aAge, B_name, bAge)
	fmt.Println("两人的年龄和为：", aAge+bAge)
	fmt.Println(a, b, c, d)
	fmt.Println(AA + BB + CC)

	if CC < 10 {
		println("真的")
		CC++
	}
	println(CC)

	if CC != 6 && CC < 10 {
		println("CC 不能与5")
		CC = 6

	} else {
		println(CC)

	}
	var FF int = 55
	if FF == 100 {
		FF += 20
		if FF < 100 {
			fmt.Println(FF)
		}
	}
	var i, j int
	for i = 1; i <= 9; i++ {

		for j = 1; j <= i; j++ {
			fmt.Printf("%d%s%d%s%-2d ", j, "*", i, "=", i*j)
			//fmt.Print("")

		}
		fmt.Printf("\n")
	}

	// 变量声明
	// var 关键字
	var a3 int             //   声明变量  类型 int   默认 是0
	var AB = 3             // 变量初始化 去掉数据类型  go 语言通过值自动匹配 数据类型
	var CC2 int = 200      //  声明一个变量 并且初始化一个值
	var DD = "Holle World" //  字符串
	//短声明  :=   自动判断变量 不用声明

	FF1 := "test"
	GG := 123436
	LL, kk := 990, 65 //短声明  多个变量
	yyb, uu := "text", 345

	//声明多个变量
	var x, y = 23, 46 // 也可以 声明类型 var x,y int = 2344,566
	var tt, pp = 789, "这个是多类型测试"

	var ( //括号声明多个变量
		nn int = 43
		hh     = 34
	)
	hh = 64
	// 常量  const  表示固定的值 ，在程序运行时不会别修改
	// 定义一个常量(constant)  使用const  关键字   常量定义的时候就要赋值

	const KKKK = 234234
	// KKKK = 435   // 常量定义会 不能被修改
	//多个常量定义

	//const ( // 常量可以当枚举使用
	//	AAA = 1
	//	BBB = 2
	//	CCC = 3
	//)

	// iota  常量计数器

	const ( // 常量可以当枚举使用
		AAA = 3 * iota // 从0 开始递增  步长为 3     默认步长 为1    AAA = iota
		BBB
		CCC
	)

	fmt.Println(AB, a3)
	fmt.Printf("CC2=%d,类型是%T\n", CC2, CC2)
	fmt.Printf("DD=%s,类型是%T\n", DD, DD)
	fmt.Printf("FF1=%s,类型是%T\n", FF1, FF1)
	fmt.Printf("GG=%d,类型是%T\n", GG, GG)
	fmt.Println(x + y) // 输出运算
	fmt.Println(x, y)
	fmt.Println(tt, pp) // 同时声明两个类型不同的变量
	fmt.Println(nn - hh)
	fmt.Println(LL + kk)
	fmt.Println(yyb, uu)
	fmt.Println(KKKK)
	fmt.Println(AAA, BBB, CCC)
}
