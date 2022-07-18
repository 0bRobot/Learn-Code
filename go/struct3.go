package main

import (
	"encoding/json"
	"fmt"
)

//定义结构体
//结构体标签
type User struct { // 结构体定义标签
	Id     int    `json:"id"` // json 标签
	Name   string `json:"name"`
	Avatar string `json:"avatar"` // 私有成员不能序列化    小写
}

func main() {
	fmt.Println("------结构体------------")

	// 初始化结构体
	user := new(User) // 使用new函数   指针类型
	user.Id = 231
	user.Name = "小白"
	user.Avatar = "https://blog.8kon.com"

	// 结构 序列化  成json 格式 供前端使用
	// 使用 go语言 内置函数   json.Marshal()

	Bsize, Err1 := json.Marshal(user) //序列化的对象   两个返回值  大小  error

	if Err1 != nil {
		fmt.Println("Error")
		return

	}
	fmt.Printf("%s\n", Bsize) // 输出序列化   {"id":231,"name":"小白","avatar":"https://blog.8kon.com"}

	// 反序列化
	// 定义  把内容 定义到 变量 jsonStr
	jsonStr := `{"id":22,"name":"行尾","avatar":"https://www.Mrcn.com"}` // 使用反引号

	//反序列化 需要 json.Unmarshall()

	// 接受反序列化结构体
	user2 := User{}

	Err2 := json.Unmarshal([]byte(jsonStr), &user2) // 两个参数   jsonStr 需要转成字节数组   1 传入的数据  2 传入的函数 指针类型    一个错误返回
	if Err2 != nil {
		fmt.Println("HCerror")
		return
	}
	fmt.Printf("%#v\n", user2)
	fmt.Printf("%d,%s,%s", user2.Id, user2.Name, user.Avatar) // 取出每个值
}
