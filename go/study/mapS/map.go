package mapS

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用

// map 的定义 map[KeyType]ValueType
// keyType 表示键的类型   valueType 表示 键对应值的类型
//map类型的变量默认初始值为nil，需要使用make()函数来分配内存。

//map 的定义

func DefMap() {
	c := make(map[int]string, 10)
	// 10 表示 map 的 容量 cap  其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量
	fmt.Printf("%T\n", c) //  %T  输出类型

	// map 的基本使用
	idName := make(map[int]string, 20)
	idName[0] = "张三"
	idName[1] = "tom"
	idName[2] = "peter"
	fmt.Println(idName)
	fmt.Println(idName[0])
	fmt.Printf("idNmae Type %T\n", idName)

	// map 中判断某个键值是否存在  value, ok := map[key]
	v1, ok := idName[0]
	if ok {
		fmt.Println(v1)
	} else {
		fmt.Println("该id不存在")
	}
}

// map 的遍历

func RangMap() {
	maps := make(map[string]string)
	maps["张三"] = "地址是：北京市朝阳区"
	maps["王武"] = "地址是：上海市浦东区"
	maps["李四"] = "地址是：北京市海淀区"
	maps["王美"] = "地址是：北京市顺义区"
	maps["李涛"] = "地址是：北京市朝阳区"

	for k, v := range maps {
		fmt.Println(k, v) //遍历map时的元素顺序与添加键值对的顺序无关
	}
	// 如果只想遍历 key
	for k := range maps {
		fmt.Println(k)
	}

}

//使用delete()函数删除键值对
// delete(map, key)

func DelMap() {
	maps := make(map[int]string)
	maps[0] = "test1"
	maps[1] = "test2"
	maps[2] = "test3"
	fmt.Println("删除前", maps)
	delete(maps, 0) // 在变量maps 中删除 键为0 的键值对
	fmt.Println("使用 delete 函数 删除后", maps)
	for k, v := range maps {
		fmt.Println(k, v)
	}
}

// 按照指定顺序遍历map

func SortMap() {
	rand.Seed(time.Now().UnixNano()) // Go 语言中用于设置伪随机数生成器的种子（seed）的常用方法
	smap := make(map[string]int)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		smap[key] = value                // 把生成的随机数 和 stu开头的值 赋值给map
	}

	//取出map中的所有key存入切片ss
	ss := make([]string, 0, 200)
	for key := range smap {
		ss = append(ss, key)
	}
	// 对切片进行排序
	sort.Strings(ss)

	// 按照排序后的key  遍历map
	for _, v := range ss {
		fmt.Println(v, smap[v])
	}

}

// 元素为map 类型的切片      切片中的 类型为map

func SliceMap() {
	s1 := make([]map[string]string, 5, 10) // s1 为 切片  切片的元素类似为 map[int]string   	切片的长度为 10  容量为20
	fmt.Printf("%T\n", s1)
	for index, v := range s1 { // 遍历切片
		fmt.Println(index, v)
	}

	// 对 切片元素中的map 进行初始化
	s1[0] = make(map[string]string, 10) //容量为10
	s1[0]["name"] = "张三"
	s1[0]["password"] = "123456.123"
	s1[0]["addr"] = "北京"
	// 遍历切片
	for index, v := range s1 {
		fmt.Println(index, v)
	}
}

//值为切片类型的map

func MapSlice() {
	m := make(map[string][]string, 5) // m 是一个 map，键为 string 类型，值为 string 类型的切片 []string
	fmt.Printf("%T\n", m)
	key := "name"       // 定义一个键名为 "name"
	value, ok := m[key] // 尝试获取键为 "name" 的值和一个表示是否存在的布尔值
	if !ok {            // 如果键 "name" 存在（ok 为 true），则获取其对应的值 value。
		// 否则（ok 为 false），创建一个空的字符串切片，并将其赋值给 value 变量。然后，在 value 的基础上，
		//使用 append 函数将字符串 "张三"、"里斯" 和 "王五" 添加到切片中。
		value = make([]string, 0)
	}
	value = append(value, "张三", "里斯", "王五")
	m[key] = value // 将更新后的值重新存回 map 中
	fmt.Println(m)
}
func MapTest() {
	ss := make(map[string]string)
	ss["123"] = "321"
	xx, x := ss["123"] // 当你从 map 中访问一个键时，会返回两个值： 与该键关联的值（如果该键存在于 map 中）。 一个布尔值，用于指示该键是否存在于 map 中。
	fmt.Println(x, xx)
}
