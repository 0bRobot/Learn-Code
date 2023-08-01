package mapS

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

//1. **学生成绩查询：**
//请创建一个存储学生姓名和对应分数的 map。然后编写一个函数，接受学生姓名作为参数，并返回该学生的分数。如果学生姓名不存在于 map 中，函数应该返回一个合适的提示信息。

func StuMap() {
	var input string
	fmt.Println("请输入学生姓名，查询成绩")

	if _, err := fmt.Scan(&input); err != nil {
		fmt.Println("输入错误")
		return
	}
	c := stuMap(input)
	if c == -1 {
		fmt.Println("该学生姓名输入错误/或信息不存在")
	} else {
		fmt.Println("改学生的成绩是:", c)
	}
}
func stuMap(name string) (z int) {
	stu := make(map[string]int)
	stu["没名字"] = 91
	stu["代号5"] = 89
	stu["随老板"] = 79
	stu["张涛"] = 80
	stu["李敏"] = 54
	stu["王五"] = 99
	v, ok := stu[name]
	if ok {
		return v

	} else {
		return -1
	}
}

//2. **删除重复项：**
//编写一个函数，接受一个字符串切片作为输入，并返回一个新的切片，其中不包含重复的元素。

func DelRep() {
	fmt.Println("请输入多个字符,已空格分割。")
	var inputStr []string                 // 创建一个字符串切片，用于存储用户输入的多个字符
	scanner := bufio.NewScanner(os.Stdin) // 创建一个标准输入的扫描器，用于读取用户的输入  bufio 是一个缓冲读写库，它提供了一些高效的方法来处理 I/O 操作，尤其是在处理大量数据时更加高效。
	for scanner.Scan() {                  // 逐行读取用户输入，直到输入为空行
		input := scanner.Text() // 获取当前行的用户输入
		if input == "" {        // 如果用户输入为空行，表示输入结束，退出循环
			break
		}
		parts := strings.FieldsFunc(input, func(r rune) bool { // 使用空格作为分隔符，将当前行的输入拆分成多个字符串
			//FieldsFunc 函数用于将字符串拆分为多个子字符串，并返回一个切片，切片中的元素是根据给定的分隔函数拆分得到的子字符串
			return r == ' '
		})
		inputStr = append(inputStr, parts...) // 将拆分得到的多个字符串添加到inputStr切片中
	}
	fmt.Println(repeat(inputStr))
}

func repeat(s []string) []string {
	x := s[:0]                       // 利用切片的重新切片特性，不会分配新的内存
	rep := make(map[string]struct{}) // 创建一个空结构体的map，用于记录字符串是否已经出现过

	for _, v := range s { // 遍历切片s中的每个元素，并将它们加入到map rep中，
		rep[v] = struct{}{} // 这样重复的元素只会在map中出现一次，确保去重后的唯一性。
	}
	for str := range rep { // 遍历map rep中的每个键，即去重后的唯一元素，
		x = append(x, str) // 然后将它们依次添加到切片x中，形成最终的结果。
	}
	return x
}

//3. **单词计数：**
//编写一个函数，接受一个字符串作为输入，并返回该字符串中每个单词的出现次数的 map。

func SrtCount() {
	// 提示用户输入一个字符串
	fmt.Println("请输入一个字符串:")
	var inputString string
	reader := bufio.NewReader(os.Stdin)      // bufio.NewReader(os.Stdin) 创建了一个可以从标准输入读取内容的读取器对象 reader
	inputString, _ = reader.ReadString('\n') // 而 reader.ReadString('\n') 则用于从标准输入中读取用户输入的一行字符串，直到用户按下回车键为止。

	// 调用 WordFrequencyMap 函数，得到每个单词出现次数的 map
	frequencyMap := WordFrequencyMap(inputString)

	// 打印结果
	fmt.Println("输入字符串:", inputString)
	fmt.Println("每个单词出现次数的 map:")
	for word, count := range frequencyMap {
		fmt.Printf("%s: %d\n", word, count)
	}

}

func WordFrequencyMap(input string) map[string]int {
	// 将输入字符串按空格分隔成多个单词
	words := strings.Fields(input) // strings.Fields(input) 函数会将输入的字符串 input 按空格分隔成多个子字符串，并返回一个切片（[]string）
	// 创建一个 map 用于存储单词出现次数
	wordMap := make(map[string]int)
	// 遍历每个单词，更新 map 中的计数
	for _, w := range words {
		wordMap[w]++ // 遍历 words 切片 出现一次 wordMap 的值(int类型)  自增1
	}

	return wordMap
}

//4. **电话簿管理：**
//创建一个电话簿的 map，将联系人的姓名作为键，联系人的电话号码作为值。编写一个函数，接受姓名和电话号码作为参数，并将其添加到电话簿中。然后编写另一个函数，接受姓名作为参数，并从电话簿中删除该联系人的信息。

func MenPhone() {
	menList := make(map[string]string)
	addNum("赵六", "1323321313", menList) // 当你将一个 map 作为参数传递给函数时，实际上传递的是一个指向底层数据结构的指针。这意味着函数中的 map 参数和原始 map 实际上指向了同一个底层数据结构，因此在函数中对 map 的修改会影响原始的 map。 // 当你将一个 map 作为参数传递给函数时，实际上传递的是一个指向底层数据结构的指针。这意味着函数中的 map 参数和原始 map 实际上指向了同一个底层数据结构，因此在函数中对 map 的修改会影响原始的 map。
	addNum("王五", "12334", menList)
	addNum("李四", "12334", menList)
	fmt.Println(menList)
	delMun("李四", menList)

}
func addNum(name, phone string, x map[string]string) {
	if _, ok := x[name]; ok { // if _, ok := x[name]; ok { ... } 这一行代码是用于检查 map 中是否已经存在给定的键 name 的常见 Go 语言写法
		fmt.Println("当前名称已经添加，请勿重复添加", name) // 当你从 map 中访问一个键时，会返回两个值
		return
	}
	x[name] = phone
}
func delMun(name string, x map[string]string) {
	if _, ok := x[name]; !ok {
		fmt.Println("该用户不存在 删个鸡毛")
		return
	}
	delete(x, name)
	fmt.Println("delete", x)
}

//5. **单词计数并排序：**
//编写一个函数，接受一个字符串作为输入，并返回按照单词频率降序排序的包含单词和频率的切片。

func FindWords() {
	words := make([]string, 0)
	fmt.Println("请输入多个字符，已空格分割：")
	reader := bufio.NewScanner(os.Stdin) // 使用 Scan() 方法读取输入，直到遇到换行符为止。
	if reader.Scan() {
		// 获取输入的字符串并去除前后的空格。
		inputStr := strings.TrimSpace(reader.Text())
		// 使用 strings.Split() 方法按空格拆分字符串，将单词存储到切片中。
		words = strings.Split(inputStr, " ")
	}
	sortWords(words)
}

func sortWords(words []string) { // S []string,x map[string]int
	tmp := make(map[string]int) // 统计每个单词的出现次数，并存储在 tmp 中，tmp 的键是单词，值是出现的次数
	for _, word := range words {
		tmp[word]++
	}
	fmt.Println(tmp)
	//将 tmp 中的数据转换为两个切片：wordsList 存储单词，wordsNum 存储出现次数。
	wordsList := make([]string, 0, len(tmp))
	wordsNum := make([]int, 0, len(tmp))
	for words, mun := range tmp {
		wordsList = append(wordsList, words)
		wordsNum = append(wordsNum, mun)
	}

	sort.SliceStable(wordsList, func(i, j int) bool { // 按照单词出现次数（wordsNum）对 wordsList 进行降序排序
		return wordsNum[i] > wordsNum[j]
	})
	sort.SliceStable(wordsNum, func(i, j int) bool { // 按照单词出现次数（wordsNum）对 wordsNum 进行降序排序，保持与 wordsList 对应的关系。
		return wordsNum[i] > wordsNum[j]
	})
	fmt.Println(wordsList)
	fmt.Println(wordsNum)
}

//6. **判断两个 map 是否相等：**
//编写一个函数，判断两个给定的 map 是否相等（即包含相同的键和值）。

func EquMap() {
	s := make(map[string]int)
	s1 := make(map[string]int)
	s["n"] = 1
	s["i"] = 2

	s1["n"] = 1
	s1["i"] = 2

	fmt.Println(checkMap(s, s1))

}
func checkMap(x, y map[string]int) bool {
	// 判断长度是否相等
	if len(x) != len(y) {
		return false
	}
	// 判断两个 map 中的键值对是否完全一致
	for key, val1 := range x { // 通过 range 关键字遍历 x 中的键值对。在每次迭代中，key 变量保存当前键的值，val1 变量保存当前键对应的值。
		val2, ok := y[key]       // 在 y 中查找键 key 对应的值，并将其赋值给变量val2。如果键 key 存在于 y 中，则 ok 会被设置为 true，否则为 false。
		if !ok || val1 != val2 { // 通过条件判断来检查以下两种情况： !ok：如果键 key 在 y 中不存在（ok 为 false），则 !ok 条件成立。
			//val1 != val2：如果键 key 在 y 中存在，但其值 val2 不等于在 x 中对应的值 val1，则 val1 != val2 条件成立。
			return false
		}
	}
	for key, val2 := range y { // 额外判断 y 中的键值对是否完全包含在 x 中
		val1, ok := x[key]
		if !ok || val1 != val2 {
			return false
		}
	}
	return true
}

//7. **Map 的键值对交换：**
//编写一个函数，接受一个 map，将其键和值进行交换，并返回新的 map。

func SwapMap() {
	s1 := map[string]int{"ni": 1, "test": 2}
	fmt.Println(s1)
	fmt.Println(mapSwap(s1))
	//需要注意的是，如果原始 map 中有多个键对应相同的值，由于 map 的键必须唯一，键值交换后的 map 中只会保留其中一个键值对。其他相同值对应的键会被覆盖。
}
func mapSwap(x map[string]int) map[int]string {
	newMap := make(map[int]string) // 新的 map，用于存储键值交换后的结果
	for k, v := range x {          //遍历原始 map，进行键值交换
		newMap[v] = k
	}
	return newMap
}
