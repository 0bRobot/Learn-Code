package main

import (
	"fmt"
	"strings"
)

func main() {

	// 表示一个windows下 一个文件路径
	filename := "C:\\Go\\hello\\1.exe" //  \\ 转义\G  \h  避免歧义
	fmt.Println(filename)

	//	定义多行字符串
	ss1 := `你好
世界
     666`
	fmt.Println(ss1)

	// 字符串操作
	fmt.Println(len(ss1)) // 求字符串长度
	// 字符串拼接
	fmt.Println("hello" + "测试")
	ret := fmt.Sprintf("%s,%s", "拼接", "字符串") // %s  对应后面的两个字符串 或者变量名
	// fmt.Sprintf  返回值给变量
	fmt.Println(ret)
	// 分割
	s5 := "你:是:谁"
	fmt.Println(strings.Split(s5, ":")) // strings.Split 分割  指定分割某个字符串  变量指定分隔符
	//	 包含
	fmt.Println(strings.Contains(s5, "是")) // 判断字符串是否包含 某个值  返回一个bool 值

	//前后缀判断
	fmt.Println(strings.HasPrefix(s5, "你")) //前缀包含  返回bool 值
	fmt.Println(strings.HasSuffix(s5, "谁")) // 后缀包含  返回bool 值

	// 判断字符串出现的位置
	fmt.Println(strings.Index(s5, ":"))     // 第一个 ： 出现的位置  3
	fmt.Println(strings.LastIndex(s5, ":")) // 最后一个 ： 出现的位置  7
	// join 操作 拼接
	sli := []string{"你", "我", "他"}   // 切片
	fmt.Println(strings.Join(sli, "-")) // 用指定符号 拼接字符串

}
