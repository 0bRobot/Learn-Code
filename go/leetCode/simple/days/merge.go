package days

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {

	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	fmt.Println(nums1)

}

func Exp18() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 2
	merge(nums1, m, nums2, n)
}
