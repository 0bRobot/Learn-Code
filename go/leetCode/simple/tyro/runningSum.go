package tyro

import (
	"fmt"
	"strconv"
)

// 一维数组的动态和
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ { // 遍历整数切片，从第二个元素开始（索引为 1）
		nums[i] = nums[i] + nums[i-1] // 计算当前位置的元素值，等于原始值加上前一个位置的元素值
	}
	return nums // 返回修改后的整数切片，其中每个元素是累积和
}

// 将数字变成 0 的操作次数
func numberOfSteps(num int) int {
	// 初始化计数器，用于记录步骤数
	count := 0

	// 使用循环处理直到 num 变为 0
	for num > 0 {
		// 如果 num 是偶数，执行以下操作
		if num%2 == 0 {
			num /= 2 // 将 num 除以 2
		} else {
			// 如果 num 是奇数，执行以下操作
			num -= 1 // 将 num 减 1
		}

		count++ // 每完成一次操作，增加步骤计数
	}
	return count // 返回总的步骤数，即将 num 减为 0 所需的步骤数
}

// 最富有客户的资产总量
func maximumWealth(accounts [][]int) int {
	max := 0 // 初始化一个变量来跟踪最富有客户的资产总额

	for _, v := range accounts {
		total := 0
		// 遍历客户的每个账户，将金额累加到客户的资产总额中
		for _, v2 := range v {
			total += v2
		}
		if total > max { // 如果当前客户的资产总额大于之前的最大值，更新最大值
			max = total
		}
	}
	return max
}

func Exp02() {
	//a := []int{1, 2, 3, 4}
	//fmt.Println(runningSum(a))

	//fmt.Println(numberOfSteps(123))

	//m := [][]int{{1, 2, 3}, {6, 4, 6}, {3, 5, 6}}
	//fmt.Println(maximumWealth(m))

}

// Fizz Buzz

//给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，并用字符串数组 answer（下标从 1 开始）返回结果，其中：
//
//answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
//answer[i] == "Fizz" 如果 i 是 3 的倍数。
//answer[i] == "Buzz" 如果 i 是 5 的倍数。
//answer[i] == i （以字符串形式）如果上述条件全不满足。

func fizzBuzz(n int) []string {
	b := make([]string, n)
	if n < 0 {
		return nil
	}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			b[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			b[i-1] = "Fizz"
		} else if i%5 == 0 {
			b[i-1] = "Buzz"
		} else {
			b[i-1] = strconv.Itoa(i) // strconv.Itoa(i) 是 Go 语言标准库中的一个函数，用于将整数 i 转换为对应的字符串表示
		}
	}
	return b
}

func Exp03() {
	z := fizzBuzz(55)
	for i, v := range z {
		fmt.Println(i, v)
	}
}

//  链表的中间结点
// 给你单链表的头结点 head ，请你找出并返回链表的中间结点。
//如果有两个中间结点，则返回第二个中间结点。

func middleNode(head *ListNode) *ListNode {
	// 使用快慢指针法
	slow := head
	fast := head
	// 使用快慢指针，快指针一次走两步，慢指针一次走一步
	// 当快指针到达链表末尾时，慢指针就在中间
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 慢指针前进一步
		fast = fast.Next.Next // 快指针前进两步
	}

	return slow // 慢指针所指的节点就是中间节点
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Exp04() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	v := middleNode(head)
	fmt.Println(v.Val)

}

// 给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
//
//如果可以，返回 true ；否则返回 false

func canConstruct(ransomNote string, magazine string) bool {
	aa := make(map[rune]int) //哈希表 aa ，用于记录 magazine 中每个字符的出现次数
	// 统计 magazine 中每个字符的出现次数
	for _, v := range magazine {
		aa[v]++
	}
	// 逐个检查 ransomNote 中的字符是否可以构造
	for _, v := range ransomNote {
		// 检查 v 是否在 aa 中存在，并且数量大于 0
		if count, exists := aa[v]; exists && count > 0 {
			aa[v]-- // 减少字符的数量，因为我们使用了一个字符
		} else {
			return false
		}

	}

	return true
}

func Exp05() {
	fmt.Println(canConstruct("aaa", "aa"))
}
