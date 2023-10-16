package algorithm

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		// 第一个条件检查是否是负数或以0结尾的非零数。如果是负数，肯定不是回文数，因为负数的倒序不会和正序相同。如果是以0结尾的非零数，也不可能是回文数，因为回文数的正序和倒序是相同的，所以不能以0结尾，除非这个数本身是0。
		return false
	}

	a := 0
	for x > a {
		a = a*10 + x%10 // 将x的最后一位数字加到a的末尾
		//fmt.Println("一 a", a)
		//fmt.Println("一 x", x)
		x /= 10 // 去掉x的最后一位数字 x = x/10
		//fmt.Println("二 x", x)
		//a 最终存储的是 x 的倒序数值。

	}
	return x == a || x == a/10
	// 如果 x 是偶数位数，例如1221，那么 x 的倒序数值应该和原来的数值相等。
	//如果 x 是奇数位数，例如121，那么 x 的倒序数值应该比原来的数值多一位数字

}
func Exp06() {
	fmt.Println(isPalindrome(12121))

}
