package Interface

import (
	"errors"
	"fmt"
	"math"
)

// 接口的定义

//type i1 interface {
//}

// 接口的实现条件

// Go语言中实现接口的条件总结如下：
// 方法签名匹配：具体类型必须实现接口中定义的所有方法，并且方法签名必须完全匹配，包括方法的名称、参数列表和返回值类型。
// 隐式实现：接口的实现是隐式的，不需要显式声明。只要具体类型中的方法签名与接口定义的方法签名匹配，编译器会自动将该类型视为实现了接口。
// 多接口实现：一个类型可以实现多个接口，只要它实现了这些接口中的所有方法
// 接口类型变量可以持有实现类型的值：接口类型的变量可以持有实现接口的具体类型的值，使得我们可以通过接口变量来调用实现类型的方法。
// 空接口：Go语言中的空接口（interface{}）没有任何方法要求，因此任何类型都可以隐式实现空接口。这使得空接口可以用来存储任意类型的值。

// 值接收者和指针接收者

type shape interface { //定义一个形状接口
	Area0() float64
}

type rectangle struct { //定义一个矩形结构体
	High float64
	wide float64
}
type circle struct { //定义一个圆形结构体
	radius float64
}

func (r rectangle) Area0() float64 { // r 方法 计算矩形面积
	return r.wide * r.wide
}
func (c circle) Area0() float64 { //c 方法计算 圆形的面积
	return math.Pi * c.radius * c.radius
}

func InterfaceDemo() {
	r1 := rectangle{3.4, 6.6} // r1是直接使用rectangle类型创建的对象
	fmt.Printf("矩形的面积为:%2f\n", r1.Area0())

	c1 := &circle{5.6}
	fmt.Printf("三角形的面积为:%2f\n", c1.Area0())

	r2 := shape(rectangle{High: 2.34, wide: 2.2}) // r2是使用类型断言将rectangle对象转换为shape接口类型的对象。
	fmt.Printf("矩形的面积是%.2f\n", r2.Area0())

	c2 := shape(&circle{radius: 4.5})
	fmt.Printf("三角形的面积是%.2f\n", c2.Area0())
}

// Error 接口

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0 ！！")
	}
	return a / b, nil
}

func InterfaceDemo1() {
	a, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("result", a)
	}

}
