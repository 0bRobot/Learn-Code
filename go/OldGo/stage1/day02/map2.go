package main

import (
	"fmt"
)

func main() {
	//map4()
	//map5()
	map6()
}

func map4() {
	// map   slice

	//A := make([]int, 3, 5) // make 一个 slice
	//fmt.Printf("%T\n", A)
	//B := make(map[int]string) // make 一个 map[]
	//fmt.Printf("%T\n", B)

	// map类型的切片
	aSlice := make([]map[int]string, 4, 10)
	// 声明一个切片 其中元素类型为map   map[int]string   ， 切片中元素数量为4  容量为8
	//fmt.Printf("%T\n", AA)
	for i, v := range aSlice { //遍历切片 下标及元素
		fmt.Println(i, v)
	}
	//对切片中的map 元素进行初始化
	aSlice[0] = make(map[int]string) // 对切片X中 下标为0的map 进行初始化
	aSlice[0][1] = "123"             // 为切片中下标为0 的map赋值
	aSlice[0][2] = "456"             // 为切片中下标为0 的map赋值
	//aSlice[1] = make(map[int]string)  // 对切面中 下标为1的map 进行初始化
	aSlice[1][1] = "456" //切片中下标为1map元素  没有初始化 所有不能赋值 运行报错
	// 遍历切片
	for i, v := range aSlice {
		fmt.Println(i, v)
	}

}

func map5() {

	// 切片类型的map
	// 先定义一个map   value 类型为切片 容量为3
	aMap := make(map[string][]string, 3) //
	//fmt.Printf("%T", aMap)               // map[string][]string
	mKey := "中国" //定义map中的key

	v, ok := aMap[mKey] // 判断aMap中这个key
	if !ok {            //如果aMap中这个key 不存在
		v = make([]string, 0, 2) // 则对aMap 中的值  也就是切片 进行初始化
		//	map中值的类型为切片   切片长度为0，容量为2
	}
	v = append(v, "北京", "上海") // 使用append函数 对map中值中的切片 追加数据
	// key ： 中国  value  北京，上海   其中value 为切片
	// 将追加后的值 赋值到aMap 的切片中
	aMap[mKey] = v
	fmt.Println(aMap) // map[中国:[北京 上海]]      [北京 上海] 为切片

}

func map6() {
	//  map  key value  对调
	//必要条件 map 的值类型可以作为 key 且所有的 value 是唯一的
	iMap := map[string]string{
		"1": "张雪",
		"2": "赵天天",
	}
	oMap := make(map[string]string, len(iMap)) // 新建一个map 容量和原map 相同
	for k, v := range iMap {                   // 遍历原map
		oMap[v] = k //给 oMap 重新赋值  oMap["张雪"]="1"
	}
	for k, v := range oMap { // 遍历oMap
		fmt.Println(k, v)
	}

}
