package funcDome

import (
	"fmt"
)

// Go 语言中的函数是一段封装了可执行代码的独立块，可以通过函数名调用并传递参数，执行一系列操作，并返回一个结果
//函数的定义

//Go语言中定义函数使用func关键字，具体格式如下

func funcD1(x, a int, zzz string) (y int) { // FuncD1  函数名   (x  为 参数) (y 为 返回值)   参数和返回值 可以定义多个   其中 x a 为整形  zzz为字符串
	//func FuncD1(x, a int, zzz string) int {  // 这样的返回值 就需要在 函数体里声明 名称
	y = x + x // 函数体  实现指定功能的代码块
	return y  // 返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔
}

// 可变参数

func funcD2(x ...int) int { //注意：可变参数通常要作为函数的最后一个参数
	// 固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
	z := 0
	return z
}

// 返回值

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//返回值命名

func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 返回值补充

// 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。

//func someFunc(x string) []int {
//	if x == "" {
//		return nil // 没必要返回[]int{}
//	}
//	...
//}

//变量作用域

//全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量

// 局部变量又分为两种：
//1 数内部的局部变量：这是在函数内部声明的变量，它们只在函数的作用域内可见。在函数内部声明的变量只能在函数体内使用，并且在函数执行完毕后会被销毁。这样的局部变量对于函数的内部实现和逻辑非常有用，不会与其他函数或全局变量发生冲突。
//2 代码块内部的局部变量：Go 语言支持在代码块内部声明变量，包括 if、for、switch 等语句块。在这些代码块中声明的变量只在代码块的作用域内可见，代码块执行完毕后变量会被销毁。这样的局部变量对于临时存储数据或控制代码块逻辑非常有用。
// 如果局部变量和全局变量重名，优先访问局部变量。

// 函数类型与变量
//在 Go 语言中，函数类型是一种特殊的类型，它表示了具有特定参数和返回值的函数的类型签名。函数类型可以用于声明变量、作为函数参数、返回值等。通过定义函数类型，可以更灵活地操作函数，使代码更加模块化和可复用
// 定义 函数类型

//使用type关键字来定义一个函数类型
//type calculation func(int, int) int
//上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值

type funcMy func(int, int) int // 定义一个新的函数类型 funcMy，该类型表示接受两个 int 类型参数并返回一个 int 类型结果的函数的类型

func FuncDo() { //  add 函数实现了 funcMy 类型的函数签名，接受两个 int 类型参数并返回它们的和
	// 函数签名是函数的类型和参数列表组合的抽象描述。它描述了函数的输入和输出，但不包括函数体内的实际执行代码。函数签名由函数的名称、参数列表和返回类型组成，用来唯一标识一个函数的类型。
	var sum funcMy //  声明一个 funcMy 类型的变量 sum

	sum = add              //  将 add 函数赋值给 sum 变量
	fmt.Println(sum(3, 2)) //  调用 sum 变量，实际上是调用 add 函数，传递参数 3 和 2，并输出结果
}
func add(x, y int) int {
	return x + y
}

// 高阶函数
//高阶函数分为函数作为参数和函数作为返回值两部分。

//函数作为参数

func add1(x, y int) int {
	return x + y
}
func subtract(x, y int) int {
	return x - y
}

// 定义了一个高阶函数apply
func apply(f func(int, int) int, x, y int) int { // 它接受一个函数 f 和两个整数 x、y  然后将 x 和 y 作为参数传递给函数 f 并返回结果
	return f(x, y)
}

func FuncDom3() {
	s1 := apply(add1, 3, 4)
	s2 := apply(subtract, 1, 6)
	fmt.Println(s1)
	fmt.Println(s2)
}

// 函数组合

// 定义一个函数 double，接受一个 int 类型参数 x，返回 x 的两倍
func double(x int) int {
	return x * 2
}

// 定义一个函数 addOne，接受一个 int 类型参数 x，返回 x 加一的结果
func addOne(x int) int {
	return x + 1
}

// 定义一个高阶函数 compose，接受两个函数 f 和 g，它们都接受一个 int 参数并返回 int，返回一个新函数
func compose(a, d func(int) int) func(int) int {
	// 新函数接受一个 int 参数 x，先将 x 传递给函数 g，然后将 g 的结果传递给函数 f，最终返回 f(g(x)) 的结果
	return func(x int) int {
		return a(d(x))
	}
}

func FuncDom4() {
	//使用 compose 函数将 addOne 和 double 函数组合成一个新函数 doubleAndAddOne
	doubleAndAddOne := compose(addOne, double)
	result := doubleAndAddOne(3)
	fmt.Println(result)
}

// 函数作为返回值

// 定义一个工厂函数 makeMu，接受一个整数参数 factor
// 返回一个函数，这个函数会将传入的参数与 factor 相乘并返回结果
func makeMu(factor int) func(int) int {
	// 返回一个闭包函数，接受一个整数参数 x，
	// 将 x 与 factor 相乘并返回结果。
	return func(x int) int {
		return x * factor
	}
}

func FuncDom5() {
	// 使用 makeMu 函数生成一个乘法因子为 2 的函数 double1。
	double1 := makeMu(2)
	// 使用 makeMu 函数生成一个乘法因子为 3 的函数 triple
	triple := makeMu(3)

	// 调用 double1 函数，传入参数 5，得到 5 * 2 = 10。
	fmt.Println("double1", double1(5))
	// 调用 triple 函数，传入参数 5，得到 5 * 3 = 15。
	fmt.Println("triple", triple(5))
}

//匿名函数

//函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。
//匿名函数就是没有函数名的函数
//func(参数)(返回值){
//函数体
//}
// 匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

func FuncDom6() {
	add := func(a, b int) int { // 定义一个匿名函数 add，接受两个 int 类型参数 a 和 b，返回它们的和
		return a + b
	}
	result := add(3, 5) // 调用匿名函数 add，传入参数 3 和 5，计算并输出结果
	fmt.Println(result)

	// 定义一个匿名函数 double3，接受两个 int 类型参数 a 和 b，返回它们的积
	// 在定义同时，直接调用该匿名函数并传入参数 3 和 4，计算并输出结果
	double3 := func(a, b int) int {
		return a * b
	}(3, 4)
	fmt.Println(double3)
	// 定义一个匿名函数，接受两个 int 类型参数 a 和 b，输出它们的差
	func(a, b int) {
		fmt.Println(a - b)
	}(6, 2)
}

//匿名函数多用于实现回调函数和闭包。

// 闭包
//闭包指的是一个函数和与其相关的引用环境组合而成的实体
//闭包的实现，它使得函数内部可以持续访问外部作用域的变量。
// 闭包在很多情况下可以用来管理状态，实现私有变量，以及创建可重用和可组合的代码块。它是函数式编程中强大的概念之一

func FuncDom7() {
	a := outer() // 创建一个闭包函数，该函数引用了外部函数 outer 的局部变量
	//调用闭包函数多次，count 的值会被保持并逐渐递增
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
func outer() func() int {
	count := 0 // count 是外部函数 outer 的局部变量
	// 返回一个匿名函数, 这个匿名函数可以访问外部函数   outer 的局部变量
	return func() int {
		count++
		return count
	}
}

// defer 语句
// Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。

func calc2(s1 string, a, b int) int {
	ret := a + b
	fmt.Println(s1, a, b, ret)
	return ret
}
func DeferDome() {
	//fmt.Println("start")
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//fmt.Println("end")
	// 由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。
	x := 1
	y := 2
	defer calc2("AA", x, y) // 先进后出
	x = 10
	defer calc2("BB", x, y)
	y = 20
}

// panic / recover

func PanicDome() {
	fmt.Println("start")
	exp()

}
func exp() {
	fmt.Println("start Exp")
	panic("something went wrong  panic") // 触发panic
	fmt.Println("End  Exp")              // 这行不会执行
}

func RecoverDome() {
	// 在这里使用 defer + recover 来捕获 panic
	defer func() {
		// 使用 recover 捕获 panic，如果发生 panic，会恢复并继续执行后续代码
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()
	fmt.Println("start ")
	// 触发 panic
	panic("触发 panic")
	fmt.Println("end Main") // 这一行不会执行
}
