package days

import "fmt"

func Exp12() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	v := 2
	fmt.Println(removeElement1(nums, v))
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}

	fmt.Println(nums)
	return i
}

func removeElement1(nums []int, val int) int {
	l, r := 0, len(nums)
	for l < r {
		if nums[l] == val {
			nums[l] = nums[r-1]
			r--
		} else {
			l++
		}

	}
	fmt.Println(nums)
	return l
}
