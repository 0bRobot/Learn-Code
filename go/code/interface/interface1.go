package _interface

import (
	"fmt"
	"math"
)

// 接口的定义

type Shape interface { //定义一个名为 MyCount 的接口
	Area() float32
	Difference() bool
}

type Rectangle struct { // 定义一个矩形结构体
	Width  float32
	Height float32
}

func (r Rectangle) Area() float32 { //Area() 方法是定义在 Rectangle 结构体上的方法，用于计算矩形的面积
	return r.Width * r.Height
}
func (r1 Rectangle) Difference() bool { // 接口的 Difference 方法，用于判断是否为正方形
	return r1.Width == r1.Height
}

func InterfaceDemo() {
	s := Rectangle{ //创建一个结构体实例
		Width:  10.2,
		Height: 3.5,
	}
	// 	fmt.Println(s.Area())  直接调用计算面积方法
	var x Shape
	x = s                       // 将 将结果体变量 赋值给接口变量
	fmt.Println(x.Area())       // 通过接口调用计算面积方法
	fmt.Println(x.Difference()) //通过接口调用判断是否为正方形方法并打印
}

//接口的嵌套

type Geometry interface { // 定义了一个 Geometry 接口，它嵌套了两个小的接口 mj（面积）和 zc（周长）。
	// 通过这种方式，你在 Geometry 接口中集合了两个小接口的方法，从而可以通过一个更大的接口来调用这两个小接口中的方法。
	mj
	zc
}

type mj interface {
	Area1() float32
}
type zc interface {
	Perimeter() float32
}

type R1 struct { //定义矩形结果体
	Width  float32
	Height float32
}

// Area1 实现 方法
func (r2 R1) Area1() float32 { //方法 计算面积
	return r2.Width * r2.Height
}

// 同一个结构体多个方法时  接收者变量和接收者类型在同一个方法集中必须是一致的

// Perimeter 实现方法
func (r2 R1) Perimeter() float32 { // 计算周长
	return 2*r2.Width + 2*r2.Height
}

func InterfaceDemo1() {
	rect := R1{Width: 32.1, Height: 34.4}
	PrintGeometryInfo(rect)
	// PrintGeometryInfo 函数接受的是一个接口类型作为参数，因此你可以直接传入一个实现了接口的结构体变量，而不需要显式地创建一个接口变量并赋值。

}
func PrintGeometryInfo(geo Geometry) {
	fmt.Printf("面积 ：%.2f\n", geo.Area1())
	fmt.Printf("周长 ：%.2f\n", geo.Perimeter())
}

//接口的组合

type Reader interface { // Reader 接口定义了读取数据的能力
	Read() string
}
type Writer interface { // Writer 接口定义了写入数据的能力
	Write(date string)
}

type ReaderWrite interface { // ReaderWrite 组合了 Reader 和 Writer 接口
	Reader
	Writer
}
type Device struct {
	data string
}

func (d Device) Read() string { // Read 方法实现了 Reader 接口的 Read 方法
	return d.data
}
func (d *Device) Write(data string) { // Write 方法实现了 Writer 接口的 Write 方法
	d.data = data
}
func InterfaceDemo2() {
	d := &Device{}         // 创建一个 Device 实例的指针
	var rw ReaderWrite = d // 将 Device 实例赋值给 ReaderWrite 接口类型的变量 rw
	// 先声明一个变量 rw，并指定它的类型为 ReaderWrite 接口。接着，我们把一个 Device 类型的实例 d 赋值给这个变量 rw.
	//d 是一个 Device 类型的实例，但由于 Device 类型实现了 ReaderWrite 接口，所以可以将其赋值给 ReaderWrite 类型的变量 rw。这样，rw 变量就可以用来调用 Reader 和 Writer 接口中的方法。
	rw.Write("Hello  interface Composition") //   调用接口中的 Write 方法
	fmt.Println(rw.Read())                   // 调用接口中的 Read 方法，打印数据
}

// 空接口

func InterfaceDemo3() {
	demo01(123)
	demo01("name") // 调用 demo01 函数，传递字符串参数
	demo01(3.14)   // 调用 demo01 函数，传递浮点数参数
	s := []string{"张三", "李四"}
	demo01(s) // 调用 demo01 函数，传递切片参数
	m := make(map[string]string)
	m["生日"] = "20220802"
	demo01(&m) //调用 demo01 函数，传递指向映射的指针参数
}
func demo01(v interface{}) { // demo01 接受一个空接口参数并打印其值
	fmt.Printf("%v\n", v)
}

// 空接口2

func InterfaceDemo4() {
	data := make(map[string]interface{})
	//添加不同的值到map 中
	data["姓名"] = "张三"
	data["年龄"] = 24
	data["成绩"] = []int{67, 77, 97, 88}
	data["结婚"] = false

	for i, v := range data {
		fmt.Println(i, v)
	}

	dd := make(map[interface{}]interface{})
	dd[1] = "id"
	dd["年龄"] = 34
	dd["北京市朝阳区"] = "地址"
	dd[false] = "未婚"
	dd[2] = data

	for i, v := range dd {
		fmt.Println(i, v)
	}
}

// demo

type MyCount interface { // 定义一个名为 MyCount 的接口
	Area0() float64
	Perimeter() float64
}

// Circle 圆形结构体
type Circle struct {
	Radius float64 // 半径
}

// Rectangle0 矩形结构体
type Rectangle0 struct {
	Width  float64
	Height float64
}

func (c Circle) Perimeter() float64 { // 在 Circle 结构体上实现 Perimeter 方法，计算圆的周长
	return c.Radius * 2 * math.Pi // 定义园周长的计算方法
}
func (c Circle) Area0() float64 { // 在 Circle 结构体上实现 Area0 方法，计算圆的面积
	return c.Radius * c.Radius * math.Pi // 定义园面积的计算方法
}

func (e Rectangle0) Perimeter() float64 { // 在 Rectangle0 结构体上实现 Perimeter 方法，计算矩形的周长
	return 2*e.Width + 2*e.Height
}

func (e Rectangle0) Area0() float64 { // 在 Rectangle0 结构体上实现 Area0 方法，计算矩形的面积
	return e.Width * e.Height
}

func InterfaceDemo5() {
	s := Rectangle0{3.1, 5.4} // 创建一个矩形实例
	y := Circle{4}            // 创建一个圆形实例
	x := []MyCount{s, y}      // 将矩形和圆形实例放入 MyCount 接口切片中   x 是一个接口类型的切片

	for _, MyCount := range x { // 遍历切片中的每个实例
		fmt.Printf("%T\n", MyCount) // 打印当前实例的类型
		fmt.Printf("Area %.2f\n", MyCount.Area0())
		fmt.Printf("Perimter %.2f\n", MyCount.Perimeter())
		fmt.Println()
	}

}
