package days

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] { // 对于每个遍历到的元素，判断其是否与前一个元素相等。比较 nums[fast] 和 nums[fast-1]。
			nums[slow] = nums[fast] // 如果当前元素不等于前一个元素，则将当前元素复制到新数组中，位置由 slow 索引确定，然后将 slow 自增 1。
			slow++                  // 2
		}

	} //循环结束后，slow 的值即为新数组的长度。原数组 nums 中前 slow 个元素为新数组，
	// 该算法通过双指针（slow 和 fast）实现了在原数组上去除重复元素的操作，同时记录新数组的长度

	return slow

}

func Exp10() {

	s := []int{1, 1, 1, 2}
	fmt.Println(removeDuplicates(s))
	fmt.Println(s)
}
