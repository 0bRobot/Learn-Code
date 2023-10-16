package days

import (
	"fmt"
	"sort"
)

func deleteDuplicates(head []int) []int {
	if head == nil {
		return nil
	}
	a := make([]int, 0)
	s := make(map[int]struct{})
	for _, o := range head {
		s[o] = struct{}{}
	}
	for i := range s {
		a = append(a, i)
	}

	sort.Ints(a)
	return a

}

func Exp17() {
	h := []int{1, 8, 4, 23, -1, 3, -40, 2, 3, 2, 4}
	fmt.Println(deleteDuplicates(h))

}
