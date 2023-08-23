package mangerStudents

import "fmt"

// 使用“面向对象”的思维方式编写一个学生信息管理系统。
// 学生有id、姓名、年龄、分数等信息
// 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能

// Student 存储学生信息的结构体
type Student struct {
	Id     int
	Name   string
	Age    int8
	Source int8
}

// MangerStu 管理系统 存储学生信息 还可以 查看学生列表，添加，编辑，删除学生
type MangerStu struct {
	StuInfo map[int]*Student // 学生信息的映射，key是学生ID，value是指向Student的指针
	// 整体结构是：StuInfo 是一个 map，以学生的 Id 作为键，对应的值是指向 Student 结构体的指针。
	//而 Student 结构体本身也包含了学生的 Id 字段，用于在 StuInfo map 中唯一标识每个学生。这种设计允许你通过学生的 Id 进行高效的增、删、改、查操作。
}

func Manger() {
	stu := MangerStu{
		StuInfo: make(map[int]*Student, 8), // 创建一个空的学生信息映射，容量为 8
	}

	for {
		fmt.Print(`
     欢迎使用某某学生管理系统 v1.0
--------------------------------------------
	菜单:
		1. 查看所有学生信息
		2. 增加学生信息
		3. 编辑学生信息
		4. 删除学生信息
		5. 退出
`)
		// 获取用户输入
		var input int
		fmt.Printf("请输入:")
		fmt.Scanln(&input)

		switch input { // 通过输入调用对应的方法
		case 1:
			stu.showAll()
		case 2:
			stu.add()
		case 3:
			stu.edit()
		case 4:
			stu.del()
		case 5:
			return
		default:
			fmt.Println("无效的输入")
		}
	}
}

//调用方法

// showAll  输出所有学生
func (m MangerStu) showAll() {
	for id, stu := range m.StuInfo {
		fmt.Printf("id:%v name:%v age:%v score %v\n", id, stu.Name, stu.Age, stu.Source)
	}
}

// add 添加学生信息
func (m MangerStu) add() {
	fmt.Printf("按照id,姓名,年龄,成绩依次输入学生信息.空格分割\n")
	var (
		id     int
		name   string
		age    int8
		source int8
	)
	fmt.Scanln(&id, &name, &age, &source) //获取输入

	if _, ok := m.StuInfo[id]; ok {
		fmt.Println("学生信息已存在")
		return
	}
	NewStu := Student{
		Id:     id,
		Name:   name,
		Age:    age,
		Source: source,
	}
	// 将新增的学生纳入管理系统
	m.StuInfo[id] = &NewStu
	fmt.Println("新学生信息添加成功")
}

// edit 编辑学生信息
func (m MangerStu) edit() {
	//获取用户输入id
	var id int
	fmt.Printf("请输入要编辑的学生id:")
	fmt.Scanln(&id)
	if _, ok := m.StuInfo[id]; !ok {
		fmt.Println("修改的信息不存在")
		return
	}
	fmt.Println("按照,姓名,年龄,成绩输入学生新的信息.空格分割")
	var (
		name   string
		age    int8
		source int8
	)
	fmt.Scanln(&name, &age, &source) //获取输入
	m.StuInfo[id].Name = name
	m.StuInfo[id].Age = age
	m.StuInfo[id].Source = source
	fmt.Println("信息修改成功")
}

// del 删除学生信息
func (m MangerStu) del() {
	fmt.Printf("请输入要编辑的学生id:")
	var id int
	fmt.Scanln(&id)
	if _, ok := m.StuInfo[id]; !ok {
		fmt.Println("删除的信息不存在")
		return
	}
	delete(m.StuInfo, id)
	fmt.Printf("%d该学生信息删除成功\n", id)
}
