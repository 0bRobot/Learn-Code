package days

import (
	"fmt"
	"strconv"
)

func Exp19() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))

}

func addBinary(a string, b string) string {
	ans := ""                    // 用于保存二进制加法的结果
	carry := 0                   // 初始化进位为
	lenA, lenB := len(a), len(b) // 获取输入字符串a和b的长度
	n := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}(lenA, lenB) // 计算a和b的长度的最大值，作为循环的次数

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0') // 将a的每一位字符转换为数字，累加到carry
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0') // 将b的每一位字符转换为数字，累加到carry
		}
		ans = strconv.Itoa(carry%2) + ans // 计算当前位的值，并添加到结果ans的最前面
		carry /= 2                        // 计算进位
	}
	if carry > 0 { // 处理可能存在的最高位的进位
		ans = "1" + ans
	}
	return ans
}
