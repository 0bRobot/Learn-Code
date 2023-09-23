package Interface

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

//1. **图形计算器：**
//创建一个图形计算器程序，它可以计算不同图形的面积（如圆形、矩形、三角形等）。
//定义一个接口 `Shape`，其中包括一个方法 `Area() float64`，然后实现多个结构体，每个结构体代表一个图形，并实现 `Shape` 接口。
//用户可以选择图形类型并输入必要的参数，然后程序根据用户选择的图形计算面积。

type Shape interface {
	Area() (float64, float64)
}
type Rectangle struct { // 定义矩形参数结构体
	High float64
	Wide float64
}
type Circle struct { // 定义圆参数结构体
	radius float64
}
type Triangle struct { // 定义三角形参数结构体
	base   float64
	height float64
}

func (r Rectangle) Area() (float64, float64) { // 矩形 计算面积和周长的方法实现
	return r.Wide * r.High, 2*r.High + 2*r.Wide

}
func (c Circle) Area() (float64, float64) { //圆 计算面积和周长的方法实现
	return c.radius * c.radius * math.Pi, c.radius * 2 * math.Pi
}
func (t Triangle) Area() (float64, float64) { //三角形 计算面积和周长的方法实现
	return t.base * t.height / 2, t.base + t.height + math.Sqrt(t.base*t.base+t.height*t.height)
}

func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Ex01() {
	for true {
		fmt.Println("-------------------------------------")
		fmt.Println("----------欢迎使用图形计算器---------")
		fmt.Println("----------请选择要计算的图形---------")
		fmt.Println("------------- 1. 矩 形  -------------")
		fmt.Println("------------- 2. 圆 形  -------------")
		fmt.Println("------------- 3. 三角形 -------------")
		fmt.Println("------------- 0. 退 出  -------------")
		fmt.Println("-------------------------------------")
		var choice int
		fmt.Printf("请输入: ")
		fmt.Scanln(&choice)

		var shapeS Shape

		switch choice {
		case 0:
			fmt.Println("程序退出")
			return
		case 1:
			var w, h float64
			fmt.Printf("请输入矩形的宽度: ")
			fmt.Scanln(&w)
			fmt.Printf("请输入矩形的高度: ")
			fmt.Scanln(&h)
			shapeS = Rectangle{High: h, Wide: w}
		case 2:
			var r float64
			fmt.Printf("请输入圆形的半径: ")
			fmt.Scanln(&r)
			shapeS = Circle{radius: r}
		case 3:
			var b, h float64
			fmt.Printf("请输入三角形底边长度: ")
			fmt.Scanln(&b)
			fmt.Printf("请输入三角形高度: ")
			fmt.Scanln(&h)
			shapeS = Triangle{base: b, height: h}
		default:
			fmt.Println("输入无效！！")
			return
		}
		m, z := shapeS.Area()
		fmt.Printf("图型面积的为%.2f\n", m)
		fmt.Printf("图型周长的为%.2f\n", z)
		time.Sleep(4 * time.Second)
		clearScreen()
	}
}

//2. **文件处理器：**
//编写一个文件处理器程序，它能够读取不同类型的文件（如文本文件、JSON 文件、XML 文件等）并执行特定操作。
//定义一个接口 `FileProcessor`，其中包括一个方法 `Process(filename string) error`，然后创建多个结构体，每个结构体代表一个文件类型的处理器，并实现 `FileProcessor` 接口。
//用户可以选择文件类型并指定要处理的文件，然后程序将根据用户选择的文件类型执行相应的处理操作。

type FileProcessor interface {
	Process(filename string) error
}

type fileJSON struct{}

func (j fileJSON) Process(filename string) error { //
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Println("JSON文件内容: ")
	fmt.Println(string(content))
	return nil
}

type fileXML struct{}

func (x fileXML) Process(filename string) error { // XML文件处理器
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Println("XML文件内容：")
	fmt.Println(string(content))
	return nil
}

type fileText struct{}

func (t fileText) Process(filename string) error { //文本文件处理
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Println("文本文件内容: ")
	fmt.Println(string(content))
	return nil
}

func clearCmd() {
	cmd := exec.Command("cmd", "/c", "cls")
	if runtime.GOOS == "linux" {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func fileTest(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Printf("输入的文件不存在%s\n", fileName)
		return false
	} else if err != nil {
		fmt.Printf("判断文件是出错 %v\n", err)
	} else {
		return true
	}
	return false
}
func Ex02() {
	for true {

		var fileType string
		var p FileProcessor
		var fileNmae string

		fmt.Println("=======================================")
		fmt.Println("|----------  exit  退出程序 ----------|")
		fmt.Println("|--- 请输入文件类型 text、json、xml---|")
		fmt.Println("=======================================")
		fmt.Scanln(&fileType)
		switch fileType {
		case "text":
			p = fileText{}
		case "json":
			p = fileJSON{}
		case "xml":
			p = fileJSON{}
		case "exit":
			fmt.Println("程序退出！！")
			return
		default:
			fmt.Println("输入的文件类型不支持！")
			return
		}
		clearCmd()
		fmt.Println("=======================================")
		fmt.Println("|---请输入文件名(相对路径/绝对路径)---|")
		fmt.Scanln(&fileNmae)
		fmt.Println("=======================================")

		B := fileTest(fileNmae)
		if B {
			err := p.Process(fileNmae)
			if err != nil {
				fmt.Println("处理文件是出错！！")
				return
			}
			fmt.Println("文件处理完成！！")
		}
		time.Sleep(2 * time.Second)
		clearCmd()
	}
}

//3. **电子设备控制器：**
//创建一个电子设备控制器程序，它可以控制不同类型的电子设备（如灯、风扇、电视等）。
//定义一个接口 `Device`，其中包括方法 `TurnOn()` 和 `TurnOff()`，然后实现多个结构体，每个结构体代表一种类型的电子设备，并实现 `Device` 接口。
//用户可以选择要控制的设备并执行开关操作。

type Device interface { // 定义关闭 或者打开设备
	TurnOn()
	TurnOff()
}

type fans struct {
	name string
}

func (f fans) TurnOn() {
	fmt.Printf("%s的风扇已打开\n", f.name)
}
func (f fans) TurnOff() {
	fmt.Printf("%s的风扇已关闭\n", f.name)
}

type light struct {
	name string
}

func (l light) TurnOn() {
	fmt.Printf("%s的灯已打开\n", l.name)
}
func (l light) TurnOff() {
	fmt.Printf("%s的灯已关闭\n", l.name)
}

type air struct {
	name string
}

func (a air) TurnOn() {
	fmt.Printf("%s的空调打开\n", a.name)

}
func (a air) TurnOff() {
	fmt.Printf("%s的空调已关闭\n", a.name)
}

func Ex03() {
	//定义命令行参数
	fmt.Println(`
  _____             _                  _____            _             _ 
 |  __ \           (_)                / ____|          | |           | |
 | |  | | _____   ___  ___ ___  ___  | |     ___  _ __ | |_ _ __ ___ | |
 | |  | |/ _ \ \ / / |/ __/ _ \/ __| | |    / _ \| '_ \| __| '__/ _ \| |
 | |__| |  __/\ V /| | (_|  __/\__ \ | |___| (_) | | | | |_| | | (_) | |
 |_____/ \___| \_/ |_|\___\___||___/  \_____\___/|_| |_|\__|_|  \___/|_|
                                                                        
                                 v 0.1                                  
`)
	devicePtr := flag.String("d", "", "指定要控制的设备类型(fans,light,air)")
	zonePtr := flag.String("z", "", "指定设备的位置如:(客厅,卧室,厨房,餐厅,阳台)")
	actionPtr := flag.String("a", "", "操作(on,off)")

	// 解析命令行参数
	flag.Parse()

	// 根据设备类型创建设备实例
	var devices Device
	switch *devicePtr {
	case "fans":
		devices = fans{*zonePtr}
	case "light":
		devices = light{*zonePtr}
	case "air":
		devices = air{*zonePtr}
	default:
		fmt.Println("无效的输入设备！")
		return
	}
	//执行操作
	switch *actionPtr {
	case "on":
		devices.TurnOn()
	case "off":
		devices.TurnOff()
	default:
		fmt.Println("无效的操作")
		return
	}
}

//4. **车辆管理系统：**
//编写一个简单的车辆管理系统，它能够管理不同类型的车辆（如汽车、摩托车、自行车等）。
//定义一个接口 `Vehicle`，其中包括方法 `Start()` 和 `Stop()`，然后实现多个结构体，每个结构体代表一种类型的车辆，并实现 `Vehicle` 接口。
//用户可以选择车辆类型并执行启动和停止操作。

type Vehicle interface {
	Start()
	Stop()
}
type Car struct{}

func (c Car) Start() {
	fmt.Printf("汽车启动\n")
}
func (c Car) Stop() {
	fmt.Printf("汽车停止\n")
}

type Bike struct{}

func (b Bike) Start() {
	fmt.Printf("自行车启动\n")
}
func (b Bike) Stop() {
	fmt.Printf("自行车停止\n")
}

type Motorcycle struct{}

func (m Motorcycle) Start() {
	fmt.Printf("摩托启动\n")
}
func (m Motorcycle) Stop() {
	fmt.Printf("摩托停止\n")
}

func Ex04() {
	// 创建一个车辆类型到对应结构体的映射
	vehicles := map[string]Vehicle{
		"car":        Car{},
		"bike":       Bike{},
		"motorcycle": Motorcycle{},
	}
	fmt.Println("请选择要操作的设备(car,bike,motorcycle): ")
	var choice string
	fmt.Scanln(&choice)
	choice = strings.ToLower(choice) //将用户输入转换为小写

	// 根据用户选择 从映射中获取对应结果题的实例

	vehicle, exists := vehicles[choice]
	if exists != true { // if !exists
		fmt.Println("无效的输入")
		return
	}

	// 创建一个操作名称到对应操作函数的映射
	actions := map[string]func(){
		"stop":  vehicle.Stop,
		"start": vehicle.Start,
	}
	fmt.Println("启动或停止该设备(stop/start)")
	var act string
	fmt.Scanln(&act)
	act = strings.ToLower(act)

	//根据用户选择执行相应的操作
	fn, ok := actions[act]
	if ok != false {
		fn()
	} else {
		fmt.Println("无效的输入！")
	}
}

//5. **电子邮件发送器：**
//编写一个电子邮件发送器程序，它能够发送不同类型的电子邮件（如普通文本邮件、HTML 邮件、附件邮件等）。
//定义一个接口 `EmailSender`，其中包括方法 `Send(email Email) error`，然后实现多个结构体，每个结构体代表一种类型的邮件发送方式，并实现 `EmailSender` 接口。
//用户可以选择邮件发送方式并发送电子邮件。

type EmailSender interface {
	Send(email Email) error
}

type Email struct {
	To      string
	Subject string
	Body    string
}

type TxtSend struct{}

func (t TxtSend) Send(email Email) error {
	fmt.Printf("发送普通文本邮件到: %s主题: %s", email.To, email.Subject)
	fmt.Printf("邮件内容:")
	fmt.Printf(email.Body)
	return nil
}

type HtmlSend struct{}

func (h HtmlSend) Send(email Email) error {
	fmt.Printf("发送HTML邮件到: %s主题: %s", email.To, email.Subject)
	fmt.Println("邮件内容:")
	fmt.Print(email.Body)
	return nil
}

type OtherSend struct{}

func (o OtherSend) Send(email Email) error {
	fmt.Printf("发送带附件的邮件到: %s\n主题: %s\n", email.To, email.Subject)
	fmt.Println("邮件内容:")
	fmt.Println(email.Body)
	return nil
}

func StrInput(input *string) { // StrInput 函数用于从标准输入获取用户输入，并将输入存储到指定的字符串指针参数中
	// 输入的文本以换行符结束。如果读取时发生错误，将输出错误信息。
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n') // 读取用户输入的文本，以换行符为结束标志
	if err != nil {
		fmt.Println("读取时发生错误")
		return
	}
	*input = userInput // 将读取的文本存储到传递给函数的字符串指针参数中
}

func Ex05() {
	// 创建一个映射，将用户选择的邮件发送方式与对应的 EmailSender 实例关联起来
	emailSend := map[string]EmailSender{
		"text":  TxtSend{},
		"html":  HtmlSend{},
		"other": OtherSend{},
	}
	// 提示用户选择邮件发送方式
	fmt.Println("请选择邮件的发送方式: (text,html,other)")
	var ch string
	StrInput(&ch)                               // 获取用户输入的邮件发送方式
	ch = strings.ToLower(strings.TrimSpace(ch)) // 去除空白字符，并转换为小写

	// 根据用户选择获取对应的 EmailSender 实例
	emailSends, exists := emailSend[ch]
	if exists != true {
		fmt.Println("无效的选择") // 如果选择无效，输出错误信息并返回
		return
	}

	var to string
	var subject string
	var body string

	// 获取用户输入的邮件收件人地址、主题和内容
	fmt.Printf("请输入收件人地址: ")
	StrInput(&to)
	fmt.Printf("请输入邮件主题: ")
	StrInput(&subject)
	fmt.Printf("请输入邮件内容: ")
	StrInput(&body)
	fmt.Println("----------------------------------------------------")

	// 创建 Email 实例
	email := Email{
		To:      to,
		Subject: subject,
		Body:    body,
	}
	// 调用 EmailSender 实例的 Send 方法发送邮件，并处理可能的错误
	err1 := emailSends.Send(email)
	if err1 != nil {
		fmt.Println("邮件发送失败")
		return
	} else {
		fmt.Println("----------------------------------------------------")
		fmt.Println("邮件发送成功")
	}
}
