package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	strDome26()
}

// json 序列化

type Student01 struct {
	// 学生
	Id     int
	Gender string
	Name   string
}
type class11 struct {
	// 班级
	Title    string
	Students []*Student01 // 结构体指针切片
}

func strDome26() {
	c1 := &class11{
		Title:    "101", // 定义一般班级
		Students: make([]*Student01, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu11 := &Student01{
			Name:   fmt.Sprintf("stu11%02d", i),
			Gender: "男",
			Id:     i,
		}
		c1.Students = append(c1.Students, stu11) // 把学生数据追加到 students 切片中
	}
	// json 序列化 ： 结构体 --> json 格式的字符串

	data, err := json.Marshal(c1) // c1是 这个班级  json.Marshal()  结构体装换成JSON格式的字符串
	if err != nil {
		fmt.Println("json marshal failed!!")
		return
	}
	fmt.Printf("json:%s\n", data)

	//json 反序列化 ： json 格式字符串 --> 结构体
	str := `{"Title":"101","Students":[{"Id":0,"Gender":"男","Name":"stu1100"},{"Id":1,"Gender":"男","Name":"stu1101"},{"Id":2,"Gender":"男","Name":"stu1102"},{"Id":3,"Gender":"男","Name":"stu1103"},{"Id":4,"Gender":"男","Name":"stu1104"},{"Id":5,"Gender":"男","Name":"stu1105"},{"Id":6,"Gender":"男","Name":"stu1106"},{"Id":7,"Gender":"男","Name":"stu1107"},{"Id":8,"Gender":"男","Name":"stu1108"},{"Id":9,"Gender":"男","Name":"stu1109"}]}`
	c2 := &class11{}
	err = json.Unmarshal([]byte(str), c2)
	if err != nil {
		fmt.Println("json unmarshall failed!!")
		return
	}
	fmt.Printf("%#v\n", c2)
}
