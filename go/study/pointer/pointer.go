package pointer

import (
	"fmt"
)

//Go语言中的指针不能进行偏移和运算，是安全指针
//Go语言中的指针\3个概念：指针地址、指针类型和指针取值

// 指针地址和指针类型
// 每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置
//Go语言中使用&字符放在变量前面对变量进行“取地址”操作
// Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string等。

func DPointer() {
	// 取变量指针的
	a := 10
	ptr := &a // &a  获取 变量 a 的地址地址
	fmt.Printf("a:%d,ptr:%p\n", a, ptr)
	fmt.Printf("ptr:%p,ptr:%T\n", ptr, ptr) // ptr变量存储的a变量的内存地址   ptr 的 类型为 *int
	fmt.Println(&ptr)
	fmt.Println("----------------------------------------------------------")

	// 指针取值

	x := 50
	p := &x // 获取变量x的内存地址 将指针保存到变量p中
	fmt.Printf("type p %T\n", p)
	v := *p // 根据指针取值   根据变量v中存储的内存地址 找到对应的值
	fmt.Printf("type v %T\n", v)
	fmt.Printf("value v %v\n", v)

	// 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

	//变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
	//对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	//指针变量的值是指针地址。
	//对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值
}

// 指针传值

func modify01(x int) {
	x = 100
}

func modify02(x *int) { // 函数取值
	*x = 101
}
func DPointer0() {

	a := 20
	modify01(a)
	fmt.Println(a) // 20
	fmt.Println("-----------------------------------")
	modify02(&a)   // 给函数传入一个变量内存地址
	fmt.Println(a) // 101
}

// make 和new
//1. make
//make也是用于内存分配的，区别于new，它只用于slice、map以及channel的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了
//make函数的函数签名如下：
//func make(t Type, size ...IntegerType) Type
// make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作

func MakeDome() {
	var b map[string]int        // 先进行变量声明
	b = make(map[string]int, 1) // 使用make 函数 初始化 b 变量  // 不初始化 报 panic
	b["id"] = 1                 //  然后 进行 赋值
	fmt.Println(b)
}

//2. new 是一个内置的函数，它的函数签名
//func new(Type) *Type
//Type表示类型，new函数只接受一个参数，这个参数是一个类型
//*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。

func DPointer1() {
	a := new(int)
	b := new(bool)
	fmt.Printf("type %T,%p\n", a, a)
	fmt.Println(*a)
	fmt.Printf("type %T,%p\n", b, b)
	fmt.Println(*b)
	//使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值

	// 指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值
	c := new(string) // 变量c  存放地址地址
	*c = "ni a  小朋友" // 给这个内存地址中 存放对应的值
	fmt.Println(*c)  // 对指针变量 取值操作
}

//new与make的区别

//二者都是用来做内存分配的。
//make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
//而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
