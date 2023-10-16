package days

import "fmt"

func Exp20() {
	n := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(n))

}

func singleNumber(nums []int) int {
	r := 0
	for _, v := range nums {
		r = r ^ v // r ^= v  //     任何数与 0 异或都是它本身：a ^ 0 = a  任何数与自身异或都是 0：a ^ a = 0
	}
	return r
}
