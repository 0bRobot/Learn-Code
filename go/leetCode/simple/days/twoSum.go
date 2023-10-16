package days

import "fmt"

func towSun(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func towSun1(nums []int, target int) []int {
	numMap := make(map[int]int) // 创建一个空的整数哈希表，用于存储已经遍历过的数字及其索引
	for i, num := range nums {
		complement := target - num               // 计算当前数字与目标值的差值
		if index, ok := numMap[complement]; ok { // 检查差值是否已经存在于哈希表中
			return []int{index, i} // 如果存在，说明找到了满足条件的两个数字，返回它们的索引
		}
		numMap[num] = i // 将当前数字及其索引添加到哈希表中
	}
	return nil
}

//创建一个空的整数哈希表 numMap，用于存储已经遍历过的数字以及它们的索引。
//
//使用 for 循环遍历整数数组 nums 中的每个数字。
//
//在循环内部，首先计算当前数字与目标值 target 的差值 complement。
//
//接着，检查差值 complement 是否已经存在于哈希表 numMap 中。如果存在，说明找到了满足条件的两个数字，它们的和等于目标值 target，则返回这两个数字的索引。
//
//如果差值 complement 不存在于哈希表中，就将当前数字及其索引添加到哈希表中，以便后续的查找。
//
//如果遍历完整个数组后仍然没有找到满足条件的数字组合，就返回一个空切片表示未找到

func Exp() {
	a := []int{2, 7, 11, 55}
	t := 9
	//fmt.Println(towSun(a, t))
	fmt.Println(towSun1(a, t))
}
