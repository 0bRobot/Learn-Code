package basePackage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type student struct {
	id   int
	name string
}

func BasePage() {
	//Print系列
	fmt.Print("输入什么打印什么！没有换行符...")
	fmt.Printf("格式化输出 id:%v", 101)
	fmt.Print("输入什么打印什么！又换行符！")
	fmt.Println()

	stu := student{
		id:   11,
		name: "赵刚",
	}

	fmt.Println(stu)             // {1 赵刚}
	fmt.Printf("%v\n", stu)      // {1 赵刚}
	fmt.Printf("%+v\n", stu)     // {id:1 name:赵刚}
	fmt.Printf("%#v\n", stu)     // basePackage.test{id:1, name:"赵刚"
	fmt.Printf("%%%v\n", stu.id) // %1  前两个 %% 表示一个%  后面的%v值的默认格式表示

	// %s
	var b = []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(b)        // [104 101 108 108 111]
	fmt.Printf("%s\n", b) // hello

	bs, err := json.Marshal(stu)
	fmt.Printf("bs:%s err:%v\n", bs, err) //%s 格式化字符串可以用于输出 []byte，会将字节切片中的内容作为字符串来处理

	fmt.Println(&stu)
	fmt.Printf("%p\n", &stu)

	// 跨度标识符
	n := 33.1415
	fmt.Printf("%f\n", n)    //33.141500  默认宽度 精度
	fmt.Printf("%9f\n", n)   //33.141500  宽度9 默认精度
	fmt.Printf("%.2f\n", n)  //33.14      默认宽度 精度2
	fmt.Printf("%9.2f\n", n) //    33.14  宽度9 精度2
	fmt.Printf("%9.f\n", n)  //       33  宽度9 精度0

	//其他flag

	s := "你好"
	fmt.Printf("%s\n", s)
	fmt.Printf("%5s\n", s)
	fmt.Printf("%-5s\n", s)
	fmt.Printf("%5.7s\n", s)
	fmt.Printf("%-5.7s\n", s)
	fmt.Printf("%5.2s\n", s)
	fmt.Printf("%05s\n", s)

	//n, m := fmt.Println("abcd") // abcd
	//fmt.Println(n, m)           // 5 <nil>

	//n, m := fmt.Print("abcd")
	//fmt.Println(n, m) // abcd4 <nil>

	//n, m := fmt.Printf("abcd\n")   // abcd
	//fmt.Println(n, m)   // 5 <nil>

	//var name string
	//var age int
	//fmt.Scan(&name, &age)
	//fmt.Print("Enter your name and age: ", name, age)

}

func BasePage1() {
	// Fprint
	fmt.Fprintln(os.Stdout, "hello") // 这是一个 fmt 包中的函数，用于格式化输出。它会将格式化的内容写入到指定的输出流中。
	//os.Stdout：这是一个代表标准输出的文件对象。在 Go 语言中，os.Stdout 是一个预定义的变量，它代表默认的标准输出流，通常是终端窗口。你可以将它传递给输出函数，以便将内容输出到终端。
	// "hello"：这是要输出的字符串内容，即要在终端上显示的文本。

	f, err := os.OpenFile("234.txt", os.O_CREATE|os.O_APPEND, 0644) // 打开名为 "234.txt" 的文件。如果文件不存在，则创建它。使用 os.O_APPEND 标志来追加内容，使用文件权限 0644。
	if err != nil {                                                 // 检查是否在打开文件时发生了错误。如果有错误，将错误信息打印到标准输出。
		fmt.Println(err)
	}
	ss := "ok?"
	//fmt.Fprintln(f, "hello") // 将字符串 "hello" 和一个换行符写入之前打开的文件 f。由于文件被以追加模式打开，这将在文件的末尾追加内容。
	fmt.Fprintf(f, "这是输入的内容%s", ss)

}

func BasePage2() {
	stu := student{
		id:   11,
		name: "赵刚",
	}
	//s := "学生id"+ stu.id +"年龄"+stu.name
	// 将学生的 id 作为整数和学生的 name 作为字符串与一些固定文本拼接起来。然而，这行代码会导致编译错误，因为在 Go 语言中，不能直接将整数类型和字符串类型相加而进行拼接，你需要使用 fmt.Sprintf 或其他字符串拼接的方法。
	s1 := fmt.Sprintf("学生id%d,学生姓名%s", stu.id, stu.name)
	// 使用 fmt.Sprintf 函数将学生的信息格式化为一个字符串。%d 是一个占位符，表示整数，%s 是一个占位符，表示字符串。这些占位符会被后面提供的值 stu.id 和 stu.name 替换。这样，格式化后的字符串会包含学生的 id 和 name
	fmt.Println(s1) //学生id11,学生姓名赵刚
}

// BasePage3  bufio.NewReader

func BasePage3() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n') //指定读到什么位置 停止
	line = strings.TrimSpace(line)       // 去掉首位的空格
	fmt.Println(err)
	fmt.Println(line)
}

// BasePage7 获取输入
func BasePage7() {
	var (
		name    string
		age     int
		married bool
	)
	// 使用竖线 | 作为分隔符
	fmt.Print("Enter data (e.g. 'Alice|25|true'): ")
	fmt.Scanf("1|%s 2|%d 3|%t", &name, &age, &married) // 输入 1|mot 2|23 3|false
	fmt.Printf("Scanned result - Name: %s, Age: %d, Married: %t\n", name, age, married)

}

func BasePage8() {

	reader := bufio.NewReader(os.Stdin) // 在标准输入上创建一个新的缓冲读取器（bufio.Reader）。这将允许我们逐行读取输入。
	fmt.Print("Enter a line of text:")
	line, _ := reader.ReadString('\n')    // 使用创建的缓冲读取器，从标准输入读取文本，直到遇到换行符（\n）。读取的内容将存储在变量 line 中。这里使用下划线 _ 来忽略 ReadString 函数可能返回的错误。
	line = strings.TrimSpace(line)        // 使用 strings.TrimSpace 函数来移除字符串两端的空白字符，包括空格、制表符和换行符。这一步骤可以用来清除用户输入文本中的额外空格和换行。
	fmt.Printf("you entered:%#v\n", line) // 使用格式化字符串输出函数 Printf，将处理过的文本以格式化形式打印出来。%#v 是一个格式化动词，它会打印出包含变量类型信息的字符串表示，方便调试和查看变量的内容。
}
