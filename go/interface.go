package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

//序列化成 xml格式
type User struct { // 结构体
	Id   int    `json:"id" xml:"id,attr"`     // 标签
	Name string `json:"name" xml:"name,attr"` // xml格式
}

//序列化
type jsonParser string // 自定义类型  序列化json

type XmlParser string // 自定义类型   序列化 xml

func (P jsonParser) Serialize(V interface{}) { // 自定义方法 接收者 p    Serialize方法  序列化方法  interface{}  没有任何返回值
	if bytes, err := json.Marshal(V); err != nil { // json.marshal()  go序列化
		panic(err) // panic() 函数  抛出这个错误
	} else { // 没有发错错误 输出
		fmt.Printf("%s\n", bytes)
	}

}
func (P XmlParser) Serialize(v interface{}) {
	if bytes, err := xml.Marshal(v); err != nil { // xml.Marshal(v)  xml 反系列化
		panic(err)
	} else {
		fmt.Printf("%s\n", bytes)
	}
}

func main() {
	fmt.Println("接口")
	user := User{ // 结构体 实例化
		Id:   102,
		Name: "王王",
	}

	var p1 jsonParser // 不同的类型 输出两次
	p1.Serialize(user)

	var p2 XmlParser
	p2.Serialize(user)

	fmt.Println("定义接口之后")
	printAny(user, p1) // 一个类型 传入不同的类型进行序列化
	printAny(user, p2)

}

type tokenParser interface { // 定义接口
	// 内部为函数原型
	Serialize(v interface{})
}

func printAny(v interface{}, p tokenParser) { // 函数传入一个接口
	p.Serialize(v)
}
