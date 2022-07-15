package main

import (
	"fmt"
	"sort"
)

func main() {

	// map 的定义   map 是一种 key  value    [键值对]   键 不能重复
	//map 关键字进行定义  map[ type(键) ] type(值) {}  键的类型 不能是 map slice func
	map1 := map[int]int{
		// 键值
		1: 3,
		2: 4,
		4: 5,
	}
	map2 := map[string]int{
		"a": 323,
		"b": 213,
		"c": 134,
	}
	// map 的遍历   /取值
	fmt.Println(map1[1])
	fmt.Println(map2["b"])
	//map 的遍历 2   使用range 遍历
	for k, v := range map2 {
		println(k, v) // map 的遍历是无序的    随机输出 。。。。。。。。。。。
	}
	// map 有序遍历 取键   赋值
	SS := []string{}
	for key := range map2 {
		SS = append(SS, key)
	}
	fmt.Println(SS)
	sort.Strings(SS)

	for _, b := range SS {
		fmt.Println(map2[b])
	}
	fmt.Println(SS)

	// map 的初始化  3   make方法
	map3 := make(map[int]string) //    make (map[type(键)]type(值))
	fmt.Println(map3)

	//map3[1]="s"
	//map4 := make(map[string]int)
	//map5 := map[string]int{
	//	"a": 1,
	//	"b": 2,
	//}
	//map4 = map5
	//fmt.Println(map4)
	test1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(test1["d"], test1["c"]) // 主键 不存在输出 0  不报错

	// 判断键值是否存在
	_, OJBK := test1["f"] //  一个返回值  一个返回bool值
	println(OJBK)

	if NO, OK := test1["a"]; OK {
		fmt.Println(NO, OK)
	} else {
		fmt.Println(NO, OK)
	}
}
