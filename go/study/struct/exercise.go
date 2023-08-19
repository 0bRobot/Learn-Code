package _struct

import (
	"fmt"
	"math"
	"sort"
)

//**题目一：定义结构体**
//定义一个名为 `Person` 的结构体，包含以下字段：`FirstName`、`LastName`、`Age`，分别表示名、姓和年龄。编写代码创建一个 `Person` 结构体对象，并打印其属性。

type person01 struct {
	firstName string
	lastName  string
	age       int
}

func StrExe1() {
	s := person01{
		firstName: "Perter",
		lastName:  "kin",
		age:       20,
	}
	fmt.Printf("%v \ns type%T\n", s, s)

}

// **题目二：计算图形面积**
// 定义一个名为 `Rectangle` 的结构体，表示矩形，包含长度和宽度字段。定义一个名为 `Circle` 的结构体，表示圆形，包含半径字段。编写代码计算矩形和圆形的面积。
type rectangle struct {
	wide float64
	long float64
}
type circle struct {
	r float64
}

// 定义矩形计算的面积方法
func (r rectangle) area() float64 {
	return r.wide * r.long
}

// 定义计算圆面积 的方法

func (c circle) cArea() float64 {
	return c.r * c.r * math.Pi
}

func StrExe2() {
	r1 := rectangle{
		wide: 5.4,
		long: 3.3,
	}
	fmt.Printf("矩形的面积为：%.4f\n", r1.area())

	c := circle{
		r: 5.4,
	}
	fmt.Printf("圆的面积为：%.4f\n", c.cArea())
}

//**题目三：成绩统计**
//定义一个名为 `Student` 的结构体，包含以下字段：`Name`、`Scores`，分别表示学生姓名和成绩（使用切片）。编写代码创建一个 `Student` 结构体对象，计算其成绩的平均值。

type student1 struct { // 定义 student1 结构体，表示学生信息
	name   string
	scores []float32
}

func (s student1) average() float32 { // 定义计算平均成绩的方法
	var total float32
	for _, v := range s.scores { // 遍历成绩切片，累加成绩
		total += v
	}

	return total / float32(len(s.scores)) // 计算平均成绩并返回
}
func StrExe3() {
	s := student1{ // 创建一个 student1 结构体对象
		name:   "李明",
		scores: make([]float32, 0), // 初始化成绩切片
	}
	s.scores = append(s.scores, 90, 88, 67.5, 77.8, 82) // 向成绩切片中添加多个成绩

	fmt.Printf("%v的平均成绩是：%.2f", s.name, s.average()) // 调用 average 方法计算平均成绩，并输出结果
}

//**题目四：购物车**
//定义一个名为 `Product` 的结构体，包含以下字段：`Name`、`Price`，分别表示产品名称和价格。定义一个名为 `ShoppingCart` 的结构体，表示购物车，包含一个 `Products` 字段（使用切片），用于存储多个产品。编写代码计算购物车中所有产品的总价。

type Product struct { // Product 结构体表示一个产品，包含名称和价格   Product 结构体的字段将会在包外部可
	name  string
	price float64
}
type ShoppingCare struct { // ShoppingCare 结构体表示购物车，包含多个产品
	products []Product // 切片的字段
}

// tail 方法计算购物车中产品的总价和数量
func (s ShoppingCare) tail() (count int, total float64) {
	for _, v := range s.products { // 遍历购物车中的每个产品
		total += v.price // 累加产品价格
	}
	count = len(s.products) // 获取购物车中产品的数量
	return count, total
}

func StrExe4() {
	p := ShoppingCare{ // 创建一个购物车实例，包含多个产品
		products: []Product{ // ShoppingCare 结构体中的 products 字段，是一个切片类型，它可以存储多个 Product 结构体
			// ShoppingCare 结构体中的 products 字段是一个切片类型，这意味着你可以在这个切片中存储多个 Product 结构体的实例。
			// 方法允许你有效地组织和管理多个相关的数据  要确保可以在不同的代码部分访问和使用这些结构体和字段，你需要将结构体名称和字段名称的首字母改为大写，以使它们在包外部可见
			// 结构体中的切片字段可以用来存储多个相同类型的值，而在这个示例中，ShoppingCare 结构体中的 products 字段可以存储多个 Product 结构体的实例，从而有效地管理购物车中的产品信息。
			{name: "小米", price: 3100}, // 第一个元素，使用完整字段名进行赋值
			{"华为Meta", 6500},
			{"iphone13", 6200},
		},
	}
	count, total := p.tail() // 调用购物车的 tail 方法，计算总价和数量
	fmt.Println("购物车总价是：", total, "一共有", count, "商品")
}

//**题目五：银行账户**
//定义一个名为 `BankAccount` 的结构体，包含以下字段：`OwnerName`、`Balance`，分别表示账户持有人姓名和余额。编写代码创建一个 `BankAccount` 结构体对象，实现存款和取款的方法，并打印操作后的余额。

type BankAccount struct {
	ownerName string
	balance   float64
}

func (b *BankAccount) bankCount(action string, num float64) {
	if num > 0 {
		if action == "+" {
			b.balance += num
			fmt.Printf("+++++存钱成功+++++ \n存入%.2f元\n当前余额为：%.2f元\n", num, b.balance)

		} else if action == "-" && b.balance >= num {
			b.balance -= num
			fmt.Printf("-----存钱成功----- \n取出%.2f元\n当前余额为：%.2f元\n", num, b.balance)
		} else {
			fmt.Printf("___________________________\n你余额就这么点%.2f \n你还想取出来:%.2f \n想屁吃！\n___________________________", b.balance, num)
		}
	} else {
		fmt.Println("你想干啥 乱输入！！！")
	}
}

func StrExe5() {
	m := BankAccount{
		ownerName: "王五",
		balance:   3000,
	}
	//  +   存钱  -  取钱
	m.bankCount("+", 300)
}

//**题目六：学生成绩排名**
//定义一个名为 `Score` 的结构体，包含以下字段：`Name`、`Subject`、`Value`，分别表示学生姓名、科目和成绩。编写代码创建多个 `Score` 结构体对象，实现按照成绩降序排名，并打印排名结果。

type Score struct { // 定义 Score 结构体，包含学生姓名、科目和成绩
	name    string
	subject string
	value   float32
}

func StrExe6() {

	s := []Score{ // 创建一个切片 s 来存储多个 Score 结构体对象  // s  是个[]_struct.Score切片    类型是我自定义的Score 结构体
		{name: "张三", subject: "数学", value: 90},
		{name: "王五", subject: "数学", value: 79},
		{name: "李四", subject: "数学", value: 93},
		{name: "王涛", subject: "数学", value: 68},
	}
	//排序
	sort.Slice(s, func(i, j int) bool {
		return s[i].value > s[j].value // 先访问切片 在拿切片中的结构体的成绩  来排序切片
	})
	//fmt.Printf("%T", s)  []_struct.Score

	for rank, v := range s {
		fmt.Printf("第%d名: %s 成绩为：%2.f\n", rank+1, v.name, v.value)
	}
}

//**题目七：公司员工管理**
//定义一个名为 `Employee` 的结构体，包含以下字段：`Name`、`Age`、`Position`，分别表示员工姓名、年龄和职位。定义一个名为 `Company` 的结构体，表示公司，包含一个 `Employees` 字段（使用切片），用于存储多个员工。编写代码实现添加、删除和列出员工的方法。

type Employee struct { // 定义员工结构体
	name     string
	age      int
	position string
}
type Company struct { // 定义公司结构体
	employees []Employee
}

func (c *Company) PrintEmp() { // 输出员工信息的方法
	for _, v := range c.employees {
		fmt.Println(v)
	}
}
func (c *Company) addEmp(add []Employee) { // 添加员工的方法
	fmt.Printf("增加后的人员信息\n")
	c.employees = append(c.employees, add...)

}

func (c *Company) delEmp(del []Employee) { // 删除员工的方法
	fmt.Printf("删除后的人员信息\n")
	for _, toDel := range del { // 这个循环遍历 del 切片中的每个员工，其中 toDel 是当前迭代的员工结构体。
		for i, v := range c.employees { //这是一个循环，遍历 c.employees 切片中的每个员工。i 是索引，v 是切片中的当前员工结构体。
			if v == toDel { // 条件判断语句用于比较当前员工结构体 v 是否等于要删除的员工结构体 ToDel。
				c.employees = append(c.employees[:i], c.employees[i+1:]...) // 如果找到了要删除的员工，就将该员工从切片中移除。
				break                                                       // 一旦找到并删除了要删除的员工，立即跳出循环，因为我们已经完成了删除操作。
			}
		}
	}
}

func StrExe7() {
	p := Company{
		[]Employee{
			{name: "王五", age: 23, position: "AI工程师"},
			{name: "赵刚", age: 33, position: "算法工程师"},
			{name: "王淘", age: 43, position: "数据工程师"},
		},
	}

	if false {
		a := []Employee{ // 增加人员的信息
			{"李刚", 14, "增加岗位1"},
			{"明明", 44, "增加岗位2"}}
		p.addEmp(a) // 调用添加方法
	}

	if false {
		d := []Employee{
			{name: "赵刚", age: 33, position: "算法工程师"},
			//{name: "王五", age: 23, position: "AI工程师"},
		} // 删除人员信息
		p.delEmp(d) // 调用删除方法
	}

	p.PrintEmp() // 结果输出
}
