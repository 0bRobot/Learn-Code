package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//strDome18()
	//strDome19()
	strDome20()
}

// 结构体与json序列化

type student struct {
	// 学生
	id     int
	gender string
	name   string
}
type class struct {
	// 班级
	title    string
	students []*student
}

func strDome18() {
	c := &class{
		title:    "101",
		students: make([]*student, 0, 20),
	}
	for i := 0; i < 10; i++ {
		stu := &student{
			name:   fmt.Sprintf("stu%02d", i),
			gender: "男",
			id:     i,
		}
		c.students = append(c.students, stu)
	}

	// json 序列化 ： 结构体 --> json 格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed!!")
		return
	}
	fmt.Printf("json:%s\n", data)

	// json 反序列化 ： json 格式字符串 --> 结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshall failed!!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

//  结构体标签

type student1 struct {
	id     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func strDome19() {
	s1 := student1{
		id:     1,
		gender: "男",
		name:   "小迪吧",
	}
	data1, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!!")
		return
	}
	fmt.Printf("json str :%v\n", data1)
}

// 结构体和方法补充知识点

type Person1 struct {
	name   string
	age    int8
	dreams []string
}

//func (p *Person1) SetDreams(dreams []string) {
//	p.dreams = dreams
//}

func (p *Person1) SetDreams(dreams []string) {
	// 使用传入的slice的拷贝进行结构体赋值
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func strDome20() {
	p1 := Person1{name: "小可爱", age: 19}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	//修改 p1.dreams
	data[1] = "学习"
	fmt.Println(p1.dreams) // [吃饭 睡觉 打豆豆]  修改前 [吃饭 学习打豆豆]
	//同样的问题也存在于返回值slice和map的情况
}
