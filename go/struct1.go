package main //__自定义类与结构体
import (
	"fmt"
)

type XZZ struct { // 结构体的定义    type(关键字)  Name(结构体名称)  struct 关键字 {结构体成员  }
	// struct 语句定义一个新的数据类型，结构体中有一个或多个成员
	id   int
	name string
	age  int
}

type add1 struct {
	mobile string
	city   string
}

func main() {
	// 自定义数据类型  结构体

	// 初始化方法1   定义变量
	var test XZZ //普通变量的方式  var + 变量名+ 结构体名称
	//  test 变量名称  结构体name
	fmt.Printf("%#v", test) // %#v  打印结构体
	// main.XZZ{id:0, name:"", age:0}

	// 初始化方法2  简介语法

	k1 := XZZ{id: 01, name: "小张张", age: 32} //
	k2 := XZZ{id: 32, name: "2222222"}      // 初始化部分成员  age未初始化

	fmt.Printf("%#v\n类型是：%T\n", k1, k1)
	fmt.Printf("%#v\n类型是：%T\n", k2, k2)

	// 方法3 指针
	j1 := &XZZ{}              // 使用取地址方法
	j2 := &XZZ{3, "tes", 234} // 必须对应结构体成员
	fmt.Printf("%T\n,%#v\n", j1, j1)
	fmt.Printf("%T\n,%#v\n", j2, j2) // *main.XZZ  指定类型

	//方法4  new()   go内置函数
	o1 := new(XZZ) // 等价于  &XZZ{}
	o1.id = 1234
	o1.name = "张飞"
	//o2 := new()
	fmt.Printf("%T\n,%#v\n", o1, o1) // 返回指定类型

	fmt.Println("==============================================================")
	// 匿名机构体
	// 匿名结构体没有类型名称，无须通过type 关键字定义就可以直接使用
	// 定义
	var YYY struct { // 没有 type 关键字
		id   int
		name string
	}
	fmt.Printf("%T\n%#v\n", YYY, YYY)
	// 简介语法定义
	XZZ2 := struct {
		id   int
		name string
	}{
		id:   12, // 定义完成 初始化
		name: "test",
	}
	fmt.Println(XZZ2.id, XZZ2.name)
	fmt.Printf("%T\n%T\n", XZZ2.id, XZZ2.name)
	fmt.Printf("%#v\n%T\n", XZZ2, XZZ2) // 匿名结构体

	//结构体访问成员  使用 . 操作符  指针  变量 都可以使用 . 操作符
	YYY.id = 23
	YYY.name = "欧尼"

	fmt.Printf("YYY %T\n%#v\n", YYY, YYY)

	// 结构体的嵌套
	// 定义
	type t1 struct { // t1 机构体 成员需要增加成员
		id   int
		name string
		dz   add1 // add1 结构体 放在之前
	}

	p1 := &t1{12, "行", add1{mobile: "123424254", city: "京"}} //指针初始化  其他初始化不行

	p2 := t1{} // 嵌套结构体的访问
	p2.id = 111
	p2.name = "随便"
	p2.dz.mobile = "12348735"
	p2.dz.city = "哪都行"

	fmt.Printf("%T\n%#v\n", p1, p1)
	fmt.Printf("%T\n%#v\n", p2, p2)

	// 匿名成员的结构体
	type nmcy struct {
		int
		string // 成员名字 不确定  只声明类型， 切；类型不能重复
		float32
		// int  类型重复 报错
	}
	// 匿名成员初始化  // 很少用
	i1 := nmcy{}
	i1.string = "test"
	i1.int = 123
	i1.float32 = 234.32
}
