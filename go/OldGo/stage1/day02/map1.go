package main

import (
	"fmt"
	"sort"
)

func main() {
	//map1()
	//map2()
	map3()
}

func map1() {
	// map 的声明
	var m map[int]string  // map 必须初始化才能使用
	fmt.Println(m)        // map[]  没有初始化是不能使用的 nil
	fmt.Println(m == nil) // ture
	//m[123] = "24235"  // map 没有初始化不能使用
	//fmt.Println(m)

	// map 的初始化
	m1 := map[int]string{
		1: "常威", // 字面量初始化  key 不能重复
		2: "来福",
	}
	//m1 := map[int]string{} // 也可以不赋值
	m1[3] = "赵姑娘"
	fmt.Println(m1) // map[1:测试 2:小张张 3:赵姑娘]
	iName := m1[1]  // 取map的值给变量 iName
	fmt.Println(iName)

	// 使用make函数初始化map
	m2 := make(map[int]string, 20) // 20是map的容量  可选项
	m2[1] = "password"             // 给map 赋上对应的键值  key为1  值为password
	fmt.Println(m2)                // map[1:password]
}

func map2() {
	m3 := make(map[string]string)
	m3["id"] = "1"
	m3["name"] = "王小五"
	m3["passwd"] = "1234567890"
	m3["address"] = "北京市昌平区"
	fmt.Println(m3)

	// map  取值   v,ok
	v, ok := m3["name"] // 两个返回值， value   bool
	// ok为 true则value存在   ok为false 则value不存才
	fmt.Println(v, ok)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key value not found！")
	}

	// map 的遍历
	// 使用for  k,v  :=   range
	for k, v := range m3 {
		fmt.Println(k, v)
	}
	for k := range m3 {
		fmt.Println(k) // 只取k  value默认被丢弃 ，只取valve  for _,v range m3 {}  使用匿名变量
	}

}
func map3() {
	//	使用delete()内置函数删除一组键值对
	m := map[int]string{
		1: "123",
		2: "456",
		5: "09095",
		9: "13432",
		3: "66666",
		7: "7366",
		4: "789",
	}
	//fmt.Println(m) //删除前
	//delete(m, 3)
	//fmt.Println(m) // 删除后

	//按照指定顺序遍历map
	//for k, v := range m {
	//	//fmt.Println(k, v) // 未排序 无序输出
	//}
	// 取出map中的key  存入切片keys
	var keys = make([]int, 0, 10) // 注意切片元素数量为0
	for k := range m {
		keys = append(keys, k)
	}
	//对切片进行排序
	sort.Ints(keys) // 对切片的值进行排序
	//fmt.Println(keys)
	// 按照排序后的key遍历map
	for _, v := range keys {
		//fmt.Println(v) // 获取到切片排序后的值
		fmt.Println(v, ":", m[v])
	}

}
