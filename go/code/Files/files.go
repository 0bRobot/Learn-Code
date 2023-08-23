package Files

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Files() {
	file, err := os.Open("./a.txt") // 打开名为 "a.txt" 的文件，返回文件指针和可能的错误
	if err != nil {
		fmt.Println("open file failed!", err)
		return
	}
	defer file.Close() // 延迟关闭文件，确保资源被释放

	var tmp = make([]byte, 128) // 创建一个大小为 128 字节的字节切片 tmp，用于存储从文件中读取的数据

	n, err := file.Read(tmp) // 使用 file.Read 方法从文件中读取数据，存储到 tmp 字节切片中
	if err == io.EOF {
		fmt.Println("文件读完了") // 如果读取到文件末尾，打印提示信息
		return
	}
	if err != nil {
		fmt.Println("read file failed ,err", err) // 如果读取出错，打印错误信息
		return
	}
	fmt.Printf("读取了 %d字节数据\n", n) // 打印实际读取的字节数
	fmt.Println(string(tmp[:n]))  // 将读取的字节数据转换为字符串并打印出来

}

func Files1() {
	file, err := os.Open("1.txt") // 打开名为 "1.txt" 的文件，返回文件指针和可能的错误
	if err != nil {
		fmt.Println("open file failed,err", err)
		return
	}
	defer file.Close() // 延迟关闭文件，确保资源被释放
	//循环读取文件
	var content []byte          // 定义一个字节切片用于存储文件内容
	var tmp = make([]byte, 128) // 创建一个大小为 128 字节的临时字节切片 tmp，用于每次读取数据
	for {
		n, err := file.Read(tmp) // 使用 file.Read 方法从文件中读取数据，存储到 tmp 字节切片中
		if err == io.EOF {
			fmt.Println("文件读完了") // 如果读取到文件末尾，打印提示信息
			break
		}
		if err != nil {
			fmt.Println("文件读取错误！,err", err) // 如果读取出错，打印错误信
			return
		}
		content = append(content, tmp[:n]...) // 将本次读取的数据追加到 content 字节切片中
		fmt.Println(string(content))          //打印当前累积的内容（每次追加都会打印一次）
	}

}

// Files2 bufio 读取文件
func Files2() {
	// 打开名为 "main.go" 的文件，返回文件指针和可能的错误
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed  ,err", err)
		return
	}
	defer file.Close() // 延迟关闭文件，确保资源被释放
	// 创建一个 bufio.Reader 实例，用于逐行读取文件内容
	reader := bufio.NewReader(file)
	// 循环逐行读取文件内容
	for {
		line, err := reader.ReadString('\n') // 逐行读取，直到遇到换行符
		if err == io.EOF {
			if len(line) != 0 { // 如果遇到文件末尾，但最后一行可能没有换行符，所以要检查是否还有内容
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break // 退出循环

		}
		if err != nil {
			fmt.Println("read file failed ,err", err)
			return
		}
		fmt.Print(line) // 打印读取的行内容
	}

}

// Files3 整个文件读取  io/ioutil包的ReadFile方法能够读取完整的文件
func Files3() {
	// 使用 os.ReadFile 函数读取名为 "main.go" 的文件内容
	// 返回的 content 是一个字节切片，err 表示可能的错误
	content, err := os.ReadFile("./main.go")
	if err != nil {
		// 如果读取文件时出现错误，输出错误信息并终止程序
		fmt.Println("read file failed,err", err)
		return
	}
	// 将字节切片转换为字符串，并打印文件的内容
	fmt.Println(string(content))
}

// Files4 文件写入 Write和WriteString
func Files4() {
	fileName := "a.txt"
	//检查文件是否存在
	if _, err := os.Stat(fileName); err == nil { // 如果文件存在
		fmt.Print("文件已经存在是否要覆盖? (Y/N): ")
		var input string
		fmt.Scan(&input)                          // 读取用户输入
		bufio.NewReader(os.Stdin).ReadBytes('\n') // 清空输入缓冲区，丢弃回车符
		if input != "Y" && input != "y" {
			fmt.Println("程序退出！！") // 用户不同意覆盖
			return
		}
	}

	//创建或打开文件
	file, err := os.Create(fileName) // 创建文件，如果文件存在会被截断
	if err != nil {
		fmt.Println("文件创建失败", err)
		return
	}
	defer file.Close() // 确保在函数结束时关闭文件

	// 从键盘获取输入
	fmt.Print("输入要写入的内容 (输入结束后输入 Ctrl+Z):")
	reader := bufio.NewReader(os.Stdin) // 使用 bufio 包来从输入读取内容
	var lines []string
	for {
		line, err := reader.ReadString('\n') // 逐行读取输入
		if err != nil {
			break // 如果输入结束或出现错误，则退出循环
		}
		lines = append(lines, line) // 将每行内容追加到切片中
	}

	for _, line := range lines {
		trimmedLine := strings.Join(strings.Fields(line), " ") // 去除多余的空格
		//strings.Join(slice []string, sep string) string:
		//strings.Join 函数将一个字符串切片（slice）的元素通过指定的分隔符（sep）连接起来，形成一个新的字符串。
		//它接受一个字符串切片作为第一个参数，第二个参数是分隔符。函数返回一个新的字符串，其中切片的每个元素都被分隔符分开。
		trimmedLine = strings.TrimSpace(trimmedLine) + "\n" // 加上回车符
		//strings.TrimSpace 函数用于移除字符串开头和结尾的空格字符（包括空格、制表符和换行符等），返回一个新的字符串。
		//这在处理用户输入或文件内容时，常常用于去除不必要的空白字符，以便更好地处理数据。
		_, err := file.WriteString(trimmedLine) // 将处理后的内容写入文件
		if err != nil {
			fmt.Println("文件写入失败！", err)
			return
		}
	}
	fmt.Printf("将 %d 个字节写入文件", len(strings.Join(lines, ""))) // 打印写入的字节数
}

// Files5 当使用 bufio.NewWriter 来写入文件时，它会提供缓冲功能，可以有效地减少频繁的磁盘写入操作，从而提高写入效率
func Files5() {
	fileName := "outPut.txt"

	// 创建或打开文件
	file, err := os.Create(fileName) //os.Create 会将该文件截断（即清空文件内容）
	if err != nil {
		fmt.Println("文件创建失败", err)
		return
	}
	defer file.Close() //os.Create 会将该文件截断（即清空文件内容）

	writer := bufio.NewWriter(file) // 创建一个新的 bufio.Writer，关联到文件对象
	//通过使用 bufio.NewWriter(file)，你创建了一个新的缓冲写入器，将其关联到了文件对象 file 上。
	//意味着所有通过缓冲写入器的写入操作实际上是在将数据写入到内存缓冲区中，而不是直接写入到文件。
	//当缓冲区中的数据达到一定量，或者显式调用 Flush 方法时，数据会被刷新到底层的文件中。

	// 写入文本内容到缓冲区
	_, err = writer.WriteString("你好啊\n小可爱！！\n ")
	if err != nil {
		fmt.Println("文件写入失败")
		return
	}

	// 将缓冲区的内容写入文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("刷新缓冲区失败")
		return
	}
	fmt.Println("内容 已经写入文件")
}

func Files6() {
	str := "ok 不 ok?" // 要写入的内容

	err := os.WriteFile("./123.txt", []byte(str), 0666) // 使用 os.WriteFile 进行文件写入
	if err != nil {
		fmt.Println("文件写入失败！", err) // 如果写入过程中出现错误，打印错误信息
		return
	}

	fmt.Println("内容成功写入文件") // 在写入成功后，打印成功信息
}
