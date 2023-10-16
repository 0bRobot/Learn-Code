package days

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 { // 若数组为空，直接返回第一个位置
		return 0
	}

	insertIndex := n // 初始化插入位置为数组长度，用于处理 target 大于所有元素的情况

	left, right := 0, n-1 // 初始化左右边界

	for left <= right { // 二分查找循环
		mid := left + (right-left)/2

		if nums[mid] < target { // 计算中间位置，避免溢出
			left = mid + 1 // 目标值在右半边，调整左边界
		} else {
			insertIndex = mid // 目标值可能在左半边，更新插入位置
			right = mid - 1   // 目标值在左半边或当前位置，调整右边界
		}
	}

	return insertIndex // 返回插入位置
}

func Exp14() {
	ss := []int{1, 3, 5, 6}
	t := 2
	fmt.Println(searchInsert(ss, t))

}

func searchInsert1(nums []int, target int) int {

	n := len(nums)
	i := n
	l, r := 0, n-1 //二分查找法           0   2
	for l <= r {
		mid := (r-l)>>1 + l
		if target <= nums[mid] {
			i = mid
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	fmt.Println(nums)
	return i
}
