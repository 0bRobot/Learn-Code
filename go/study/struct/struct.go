package _struct

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func StrDome() {
	// 自定义类型
	// go 语言中 有一些基本的数据类型。  go 语言中可以使用  type 关键字 来自定义类型
	//自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义
	// type   MyInt  int
	type MyInt int16 // 自定义类型   没有 等于号
	var n MyInt
	n = 3
	fmt.Printf("值: %d  类型是:%T\n", n, n) // 值: 3  类型是:_struct.MyInt

	// 类型别名
	type MyInt2 = int16 // 类型 别名 有等于=
	var aa MyInt2
	aa = 32
	fmt.Printf("值: %d  类型是:%T\n", aa, aa) //值: 32  类型是:int16
}

//结构体

//		结构体的定义
//	 type  structName  struct 关键字

type myList struct {
	// 使用type和struct关键字来定义结构体
	FirstName, LastName string // 同一类型可以 写在同一行
	Age                 int
	Add                 string
}

func StrDome1() {
	// 结构体实例化
	var stu1 myList
	stu1.FirstName = "张" // 可以通过. 来访问 结构体的字短(成员)
	stu1.LastName = "那"
	stu1.Age = 30
	stu1.Add = "北京市 海淀区"
	fmt.Println(stu1)
	//当然 也可以 简短
	stu2 := myList{FirstName: "李", LastName: "明", Age: 19, Add: "上海市 思明区"}
	fmt.Println(stu2)
}

// 匿名结构体

func StrDome2() {

	var users struct {
		name string
		age  int
	}
	users.name = "王涛"
	users.age = 24
	fmt.Printf("值：%v,类型：%T\n", users, users) // 值：{王涛 24},类型：struct { name string; age int }

	user2 := struct {
		name, sex string
		age       int
		add       string
	}{
		name: "李梅",
		sex:  "女",
		age:  33,
		add:  "北京市",
	}
	fmt.Printf("值：%v 类型 %T\n", user2, user2) // 值：{李梅 女 33 北京市} 类型 struct { name string; sex string; age int; add string }

}

//  创建 指针类型的结构体

type person struct {
	name string
	age  int
	add  string
}

func StrDome3() {
	var p1 = new(person) //使用了 new(person) 来分配并初始化一个 person 类型的指针 p1
	// new 函数用于创建指向新分配的零值的指针。因此，p1 是一个指向 person 类型的指针，它指向一个新分配的 person 结构体实例
	p1.name = "张三"
	p1.age = 33
	p1.add = "上海市"
	fmt.Printf("值：%v,类型是：%T\n", p1, p1) // 值：&{张三 33 上海市},类型是：*_struct.person

}

//取结构体的地址实例化

func StrDome4() {
	p2 := &person{} // 直接使用 &person{} 创建了一个 person 类型的指针 p2，并对其字段进行了赋值。这里你手动创建了一个结构体实例并将指针指向它
	p2.name = "憨憨"
	p2.age = 18
	p2.add = "北京市"
	fmt.Printf("值：%v,类型： %T\n", p2, p2) //值：&{憨憨 18 北京市},类型： *_struct.person
}

// p1 和 p2 在输出中都是指针类型的结构体，但它们的创建方式有所不同：
//p1 是使用 new 函数分配并初始化一个新的 person 类型的指针，Go 会自动将其初始化为零值。
//p2 是手动创建一个 person 类型的指针，然后对其字段进行了赋值。
//在使用上，这两种方式都能达到相同的效果，即创建一个指向 person 结构体的指针。

//  结构体的初始化
// 没有初始化的结构体 ，其成员变量都是 对应其类型的零值

func StrDome5() {

	var p3 person
	fmt.Printf("%#v\n", p3) //  _struct.person{name:"", age:0, add:""}

	//使用键值 对结构体初始化时，值对应改字段的初始值
	p4 := person{
		name: "憨憨",
		age:  18,
		add:  "上海市",
	}
	fmt.Printf("值 :%#v\n", p4) // 值 :_struct.person{name:"憨憨", age:18, add:"上海市"}
	// 也可以对结构体指针进行键值对初始化
	p5 := &person{
		name: "憨憨1",
		age:  23,
		add:  "天津",
	}
	fmt.Printf("值 :%#v\n", p5) // 值 :&_struct.person{name:"憨憨1", age:23, add:"天津"}
	// 某些字段 没有初始值 可以 不写
	p6 := &person{
		name: "狗蛋",
	}
	fmt.Printf("值 :%#v\n", p6) // 值 :&_struct.person{name:"狗蛋", age:0, add:""}
	// 使用值的列表 初始化
	p7 := &person{
		"王刘",
		23,
		"天安门",
	}
	fmt.Printf("值 ：%#v\n", p7) // 值 ：&_struct.person{name:"王刘", age:23, add:"天安门"}

	//使用这种格式初始化时，需要注意：

	//必须初始化结构体的所有字段。
	//初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	//该方式不能和键值初始化方式混用

}

// 方法和接收者

func StrDome6() {
	// 空结构体 是 不占用空间的
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0
}

// Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）
// 方法定义的格式
//
//	func (接收者变量  接收者类型) 方法明(参数列表)(放回参数){
//		  函数体
//	}
//
// 接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写
// 接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型
// 方法名、参数列表、返回参数：具体格式与函数定义相同。
type per struct { // 定义一个名为 per 的结构体类型
	name string
	age  int
}

func newPer(name string, age int) *per { // 定义一个名为 newPer 的函数，用于创建并返回一个 per 类型的结构体指针
	return &per{
		name: name,
		age:  age,
	}
}
func (p per) Dream() { // 在 per 类型上定义一个名为 Dream 的方法
	fmt.Printf("%s的梦想是学好go语言\n", p.name) // 这个方法接收一个 per 类型的接收者 p，在方法内部使用接收者的 name 字段来打印人的梦想信息
}
func StrDome7() {
	p1 := newPer("张三", 43) // 调用 newPer 函数创建一个 per 类型的结构体指针，并初始化 name 和 age 字段
	p1.Dream()             // 调用 per 类型的 Dream 方法，打印姓名的梦想信息

}

// 指针类型的接收者

type perP1 struct {
	name string
	age  int
}

func (p *perP1) setAge(newAge int) { //  修改接收者指针的 age 字段  修改了原始结构体
	p.age = newAge
}

func (p perP1) setAge2(newAge2 int) { // 修改接收者的 age 字段（但是这个修改只在方法内部有效，不会影响原始值）
	p.age = newAge2
}
func StrDome8() {
	// 指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
	p1 := perP1{name: "李梅", age: 23}
	fmt.Println(p1.age)
	p1.setAge(98) // 因为使用了指针接收者，直接修改了原始结构体
	fmt.Println(p1.age)

	p2 := perP1{name: "王五", age: 55}
	fmt.Println(p2.age)
	p2.setAge2(43) // 修改了接收者的 age 字段，但修改只在方法内部有效，不会影响原始值
	fmt.Println(p2.age)
}

// 任意类型添加方法

type myInt int //  自定义类型

func (m myInt) sayHello() { // 添加了 一个自定义类型的 方法
	fmt.Println("hello,i am  int")
}
func StrDome9() {
	var m1 myInt
	m1.sayHello() //
	m1 = 100      // 由于 myInt 是一个自定义类型，它不能直接赋值给内置类型。但是，你可以进行类型转换。将 m1 赋值为 100，实际上是将 100 转换为 myInt 类型。
	fmt.Printf("%#v   %T\n", m1, m1)
}

// 结构体的匿名字段
type phone struct {
	string
	float32
	int32
}

func StrDome10() {
	xiaomi := phone{"三星屏幕", 6.7, 5000}
	// 匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
	fmt.Printf("%#v,%T\n", xiaomi, xiaomi) // _struct.phone{string:"三星屏幕", float32:6.7, int32:5000},_struct.phone
	fmt.Printf("小米手机屏幕是：%v,尺寸是：%v英寸,电池大小：%v毫安\n", xiaomi.string, xiaomi.float32, xiaomi.int32)
}

// 结构体嵌套
type address struct {
	shen string
	shi  string
}
type users struct { // 定义一个名为 users 的结构体类型，包含一个嵌套的 address 结构体
	name    string
	sex     string
	address address
}

type users1 struct { // 定义一个名为 users1 的结构体类型，嵌套 address 结构体的字段使用匿名字段
	name string
	sex  string
	address
}

func StrDome11() {
	user1 := users{
		name: "张涛",
		sex:  "男",
		address: address{
			shen: "陕西省",
			shi:  "宝鸡市",
		},
	}
	fmt.Printf("user1 %#v\n", user1) // user1 _struct.users{name:"张涛", sex:"男", address:_struct.address{shen:"陕西省", shi:"宝鸡市"}}

	var user2 users1
	user2.name = "李涛"
	user2.sex = "男"
	user2.shen = "江苏省"        // 匿名字段可以省略
	user2.address.shi = "南京市" // 匿名字段默认使用类型名作为字段名
	fmt.Printf("user2 %#v\n", user2)
	// 当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。
}

// 嵌套结构体的字段名冲突
type addr struct {
	city       string
	createTime string
}
type mail struct {
	acc        string
	createTime string
}
type useR struct {
	name string
	sex  string
	addr // 嵌套 addr 结构体，字段名被隐式继承
	mail // 嵌套 mail 结构体，字段名被隐式继承
}

func StrDome12() {
	var user3 useR
	user3.name = "李米"
	user3.sex = "女"
	//user3.acc = "24jks@gmail.com" // 输出的是最后一次赋值的值
	user3.mail.acc = "3252jif@yaho.com"
	user3.addr.createTime = "2023-08-11" // 通过 addr 结构体的字段名 createTime 访问匿名字段
	//user3.createTime="2023-08-11"   //由于存在两个匿名字段，且有相同的字段名，必须明确指定是哪个结构体的字段
	user3.mail.createTime = "2023-08-20" // 通过 mail 结构体的字段名 createTime 访问匿名字段
	fmt.Printf("user3 %#v\n", user3)
	fmt.Println(user3)
}

// 结构体的继承

type personD struct { // 基础结构体
	name string
	age  int
}
type student struct { // 继承Person结构体

	personD    // 	嵌套 personD结构体
	schoolName string
}

func StrDome13() {
	s := student{personD{ // 创建一个Student对象
		name: "张老三",
		age:  23,
	},
		"张小小"}
	// 访问继承的字段
	fmt.Println("name", s.name)
	fmt.Println("age", s.age)
	fmt.Println("schoolNmae", s.schoolName)
	// 修改继承的字段
	s.age = 43
	fmt.Println(s.age)
	// Student 结构体嵌入了 Person 结构体，从而继承了 Person 结构体中的字段。你可以通过 s.Name 和 s.Age 直接访问 Person 结构体中的字段，也可以通过 s.SchoolName 访问 Student 结构体中的字段。
}

// 结构体与JSON序列化

// JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式
// JSON键值对是用来保存JS对象的一种方式，键/值对组合中的键名写在前面并用双引号""包裹，使用冒号:分隔，然后紧接着值；多个键值之间使用英文,分隔。

type student01 struct { // 定义学生信息的结构体
	Id   int
	Sex  string
	Name string
}
type class struct { // 定义班级信息的结构体
	Title    string
	Students []*student01 // 使用切片的指针来存储学生对象的指针
	// Students 字段是一个切片的指针，指向了 student01 结构体类型的指针。尽管 Students 是一个切片的指针，但它的元素是指向 student01 结构体的指针，这允许你存储和操作 student01 类型的数据。

}

func StrDome14() {
	C := &class{ // 创建一个班级对象 C
		Title:    "101",
		Students: make([]*student01, 0, 200), // 初始化一个空的学生切片的指针，容量为 200
		// 往 C.Students 切片中添加学生对象时，你实际上是将 student01 结构体的指针添加到切片中。这样做的好处是，你可以有效地管理学生对象并避免复制大量的数据。每个学生对象本身仍然保留在内存中，只是 Students 切片中存储了这些学生对象的指针。
	}
	for i := 0; i < 3; i++ { //循环创建 3 个学生对象并添加到班级 C 中
		stu := &student01{
			Name: fmt.Sprintf("stu%02d", i),
			Sex:  "男",
			Id:   i,
		}
		C.Students = append(C.Students, stu) // 将学生对象的指针添加到班级的学生切片中
	}
	// 打印每个学生的内容

	//for _, stu := range c.students {
	//	fmt.Println(stu.name, stu.sex, stu.id)
	//}

	// json  序列化：  结构体————>  json 格式字符串
	data, err := json.Marshal(C) // 将班级 C 序列化为 JSON 格式的字符串
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	// json:{"Title":"101","Students":[{"Id":0,"Sex":"男","Name":"stu00"},{"Id":1,"Sex":"男","Name":"stu01"},{"Id":2,"Sex":"男","Name":"stu02"}]}

	//  json 反序列化   json 格式字符串  ————> 结构体
	// 准备一个 JSON 格式的字符串，包含一个班级信息和三个学生信息
	str := `{"Title":"101","Students":[{"Id":0,"Sex":"男","Name":"stu00"},{"Id":1,"Sex":"男","Name":"stu01"},{"Id":2,"Sex":"男","Name":"stu02"}]}`
	C1 := &class{}
	err = json.Unmarshal([]byte(str), C1)
	if err != nil {
		fmt.Println("json unmarshal failed!!")
		return
	}
	// 遍历并打印 C1 中每个学生的信息
	for _, stu := range C1.Students {
		fmt.Printf("Id:%v Sex:%v Nmae:%v  ", stu.Id, stu.Sex, stu.Name)
	}
}

// 结构体标签

// Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来
// 结构体tag由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔
// 注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值

type students struct {
	Id   int    `json:"Id"` // Id int 字段使用了结构体标签 json:"Id"，这表示在 JSON 序列化和反序列化时，该字段应该使用 "Id" 作为字段名称
	Name string // 默认的字段名称为结构体字段名
	Sex  string // Name string 字段和 Sex string 字段没有使用标签，它们会使用默认的字段名作为 JSON 字段名。
}

func StrDome15() {
	s1 := students{
		Id:   01,
		Name: "米玛",
		Sex:  "女",
	}
	date, err := json.Marshal(s1) // 将 students 对象序列化为 JSON 格式的字符串
	if err != nil {
		fmt.Println("json marshal  failed!")
		return
	}
	fmt.Printf("json str :%s\n", date)
}

// 其他问题

type P11 struct {
	name   string
	age    int
	dreams []string
}

func (p *P11) setDreams1(dreams []string) {
	p.dreams = dreams
}

//func (p *P11) setDreams1(dreams []string) {
//	p.dreams = make([]string, len(dreams))
//	copy(p.dreams, dreams)    // 使用 copy 函数来避免切片共享问题，从而创建一个新的底层数组，不受修改的影响。
//}

func StrDome16() {
	s := P11{name: "李四", age: 20}
	date := []string{"玩游戏", "睡觉觉", "看电影"}
	s.setDreams1(date)

	date[1] = "不睡觉"       // 修改 dreams 切片的第二个元素为 "不睡觉"
	fmt.Println(s.dreams) //  注意此时输出为 ["玩游戏" "不睡觉" "看电影"]，因为 s.dreams 和 dreams 共享同一个底层数组，所以修改 dreams 会影响到 s.dreams。
}
