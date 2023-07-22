package slice

import (
	"fmt"
	"sort"
)

//2. **查找元素**:
//   编写一个函数，接收一个整数切片和一个整数，返回该整数在切片中的索引位置。如果该整数不存在于切片中，返回 -1。
//
//3. **合并切片**:
//   编写一个函数，接收多个整数切片，将它们合并成一个新的切片并返回。注意，你可以使用 `append()` 函数将多个切片合并。
//
//4. **去重**:
//   编写一个函数，接收一个整数切片并去除其中的重复元素，返回一个新的不包含重复元素的切片。
//
//5. **切片的最大值和最小值**:
//   编写一个函数，接收一个整数切片，并返回切片中的最大值和最小值。
//
//6. **切片的元素乘积**:
//   编写一个函数，接收一个浮点数切片，并计算切片中所有元素的乘积，然后返回该乘积值。
//
//7. **二维切片操作**:
//   编写一个程序，创建一个二维整数类型的切片并初始化，然后实现以下操作：
//   - 在指定行后追加一行数据
//   - 删除指定行的数据
//   - 获取指定元素所在的行和列的索引
//

func SExec1() {
	//1. **切片基本操作**:
	//   编写一个程序，创建一个整数类型的切片并初始化，然后实现以下操作：
	//   - 在切片的末尾追加一个整数
	//   - 删除切片中的第一个元素
	//   - 在切片的指定索引位置插入一个整数
	//   - 对切片进行排序
	//   - 反转切片的元素顺序

	s := []int{5, 9, 2, 4}
	// 1.1
	fmt.Println("Append int mun ", append(s, 9))
	//1.2
	s1 := []int{5, 9, 2, 4}

	fmt.Println("Del  first index", append(s1[:0], s1[1:]...))
	//1.3
	s2 := []int{10, 20, 30, 40, 50}
	fmt.Println("index  2  add  mun", append(s2[:2], append([]int{25}, s2[2:]...)...))
	// append([]int{25}, s2[2:]...)  得到   {25  30  40  50}
	// 然后在使用 append(s2[:2],{25  30  40  50})...)

	//1.4
	s3 := []int{7, 2, 8, 9, 5}
	sort.Ints(s3) // 使用 sort.Ints() 函数对整数切片进行升序排序
	fmt.Println("生序排序", s3)
	//1.4.2
	s31 := []int{7, 2, 8, 9, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(s31)))
	//我们首先导入了 sort 包。
	//接下来，我们使用 sort.IntSlice(s31) 将切片 s31 转换为一个实现了 sort.Interface 接口的类型。这样做是因为 sort.Reverse() 函数需要接收一个实现了该接口的对象作为参数。
	//然后，我们调用 sort.Reverse() 函数，并将上述转换后的对象传递给它。该函数会返回一个新的包装过的对象，其比较规则与原始对象相反（即进行逆序排序）。
	//最后，我们再次调用 sort.Sort() 函数，并将前面得到的逆序排序对象作为参数传入。这样就可以对切片进行逆序排序。
	//所以，在你提供的示例中，通过调用链式方法和函数嵌套使用，最终实现了对整数类型切片 s31 的逆序排序。
	fmt.Println("降序排序", s31)
	//	优化
	s32 := []int{7, 2, 8, 9, 5, 4, 1}

	sort.Slice(s32, func(i, j int) bool { // 使用了 sort.Slice 函数对整数类型的切片 s31 进行排序
		return s32[i] > s32[j] //  如果我们比较索引为 i 的元素和索引为 j 的元素，然后根据结果返回一个布尔值： 如果返回 true，则表示希望将索引为 i 的元素放在索引为 j 的元素之前。 如果返回 false，则表示希望保持原有顺序或将索引为 j 的元素放在索引为 i 的元素之前。
	})

	fmt.Println("降序", s32)
}

// 反转切片的元素顺序

func RevSlice() {
	s := []int{7, 9, 6, 8, 10, 11}
	sort.Slice(s, func(i, j int) bool { // 使用了 sort.Slice 函数对整数类型的切片 s 进行排序
		return i > j
	})
	fmt.Println("反转", s)

	// 方法2
	s3 := []int{2, 4, 6, 3, 8, 10, 11, 1, 7}
	reverseSlice(s3)
	fmt.Println("反转", s3)
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//2 编写一个函数，接收一个整数切片和一个整数，返回该整数在切片中的索引位置。如果该整数不存在于切片中，返回 -1。

func FindS() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("切片是:", s)
	fmt.Println(findSlice(s, 12))

}

func findSlice(s []int, index int) int {
	for i, v := range s {
		if v == index {
			return i
		}
	}
	return -1
}

//3 编写一个函数，接收多个整数切片，将它们合并成一个新的切片并返回。注意，你可以使用 `append()` 函数将多个切片合并。

func AppSlice() {
	s := []int{1, 2, 3, 4, 5, 6}
	x := []int{7, 8, 9, 10, 11, 12}
	s2 := []int{18, 23, 17, 87}
	s3 := []int{55, 43, 124, 78, 234}
	//fmt.Println(append(s, append(x, s2...)...))
	fmt.Println(appendSlice(s, x, s2, s3))

}
func appendSlice(slice ...[]int) []int {
	x := []int{}              // 定义一个空切片
	for _, s := range slice { // 函数接受多个切片  遍历值
		x = append(x, s...)
	}
	return x
}

//4. **去重**:
//   编写一个函数，接收一个整数切片并去除其中的重复元素，返回一个新的不包含重复元素的切片。

func DeduplicationSlice() {
	s := []int{1, 2, 3, 4, 4, 5, 6, 7, 0, 4, 2, 9}
	//fmt.Println(deSlice(s))
	fmt.Println(dedSlice1(s))
}
func deSlice(s []int) []int {
	x := []int{}
	seen := make(map[int]bool)

	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			x = append(x[:len(x)], v)
		}
	}
	return x
}

func dedSlice1(s []int) []int {
	x := []int{} // 创建一个空的切片来存储去重后的结果

	for _, v1 := range s { // 遍历原始切片的每个元素
		isDup := false        // 假设当前元素为非重复元素
		for _, v := range x { // 遍历去重后的切片 x，检查当前元素是否已经存在
			if v1 == v { // 如果当前元素 v1 与任何已存在于去重后切片 x 的元素值相同（即找到了一个重复项），则将标记变量 isDup 设置为 true 并跳出内层循环。
				isDup = true
				break
			}
		}
		if !isDup { //如果标记变量 isDup 仍然是 false（即没有找到任何重复项），则表示当前元素是非重复的。因此，在外部循环迭代结束后，我们可以将它附加到新切片 x 中以保留唯一值。
			x = append(x, v1)
		}
	}
	return x
}

//   编写一个函数，接收一个整数切片，并返回切片中的最大值和最小值。

func BjSlice() {
	s := []int{1, 4, 7, 2, 8, 9, 22, 31, 0}
	//fmt.Println(bj(s))
	fmt.Println(bj2(s))
}

func bj(x []int) (min, max int) {
	min, max = x[0], x[0] //初始化最小值和最大值为切片中的第一个元素
	for _, v := range x {
		if v < min { // // 判断当前元素是否比最小值小，若是则更新最小值
			min = v
		}
		if v > max { // 判断当前元素是否比最大值大，若是则更新最大值
			max = v
		}
	}
	return min, max
}

func bj2(x []int) (min, max int) { // 方法2
	sort.Ints(x) //  使用 sort 包中的排序函数对切片进行排序
	min = x[0]   // 切片已经排序，s[0] 就是最小值，s[len(s)-1] 就是最大值
	max = x[len(x)-1]

	return min, max
}

//   编写一个函数，接收一个浮点数切片，并计算切片中所有元素的乘积，然后返回该乘积值。

func MuSlice() {
	s := []float64{2.2, 1.3, 3.1, 2.7}
	fmt.Println((mulSlice(s)))
	fmt.Println((product(s)))
}

func mulSlice(x []float64) float64 {
	res := 1.0 //  不能设为0  0 乘以任何数  都为0
	for _, v := range x {
		res = v * res // 避免出现    res = v * v  最终得到的 res 只是最后一个元素的平方
	}
	return res
}

func product(x []float64) float64 {
	if len(x) == 0 {
		return 1
	}
	res := x[0]
	for i := 1; i < len(x); i++ {
		res *= x[i]
	}
	return res
}

//   编写一个程序，创建一个二维整数类型的切片并初始化，然后实现以下操作：
//   - 在指定行后追加一行数据
//   - 删除指定行的数据
//   - 获取指定元素所在的行和列的索引

// - 在指定行后追加一行数据

func R2Slice() {
	// 在 Go 语言中，切片没有 append 方法，它只能用于对切片进行追加，而不是用于对切片中的元素进行追加
	// 要在二维切片中的指定行追加数据，需要先获取到指定行的切片，然后再向该切片中追加数据。
	s := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	s1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	n := []int{11, 22, 33, 44, 55}
	addSlice(s, 2, n)
	addSlice1(s1, 1, 100)
	//-------------------------------------------------------------------

	s2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	dSlice(s2, 0)

	s3 := [][]int{{1, 2, 3}, {4, 5, 6}, {6, 9, 2}, {7, 8, 9}}

	getSlice(s3, 2)

}

func addSlice(x [][]int, row int, date []int) { // 二维切片 指定的行 追加 切片
	if row < 0 || row >= len(x) { // 判断 row 是否合法
		fmt.Println("添加失败，行索引超出范围")
		return
	}

	rowSlice := x[row]                   // 获取指定行的切片
	rowSlice = append(rowSlice, date...) // 向切片追加数据
	x[row] = rowSlice                    //将更新后的切片重新赋值回原二维切片中的指定行
	fmt.Println(x)
}

func addSlice1(x [][]int, row int, date int) { // 二维切片 指定的行 追加 单个数字
	if row < 0 || row >= len(x) { //判断 row 是否合法
		fmt.Println("添加失败，行索引超出范围")
		return
	}
	rowSlice := x[row]                // 获取指定行的切片
	rowSlice = append(rowSlice, date) // 向切片追加数据
	x[row] = rowSlice                 //将更新后的切片重新赋值回原二维切片中的指定行
	fmt.Println(x)
}

// - 删除指定行的数据
func dSlice(x [][]int, row int) {

	if row < 0 || row >= len(x) { //判断 row 是否合法
		fmt.Println("删除行失败，超出范围")
		return
	}
	x = append(x[:row], x[row+1:]...) // 切片表达式 x[:row] 表示从 x 的开始到要删除的行（不包含该行）为止，即取前面的行部分
	//切片表达式 x[row+1:] 表示从要删除的行的下一行开始到切片x的结束，即取后面的行部分
	// 然后将这两部分拼接在一起，即可实现删除指定行的操作
	fmt.Println(x)

}

// - 获取指定元素所在的行和列的索引

func getSlice(x [][]int, value int) {
	found := false
	for i, v := range x {
		for j, v1 := range v {
			if v1 == value {
				fmt.Println(i, j)
				found = true
			}
		}
	}
	if !found {
		fmt.Println("元素不存在于切片中")
	}
}
