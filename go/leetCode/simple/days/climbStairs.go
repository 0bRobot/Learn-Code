package days

import "fmt"

func Exp16() {
	n := 4
	fmt.Println(climbStairs(n))
}

func climbStairs1(n int) int {
	if n == 0 || n == 1 { // 如果 n 为 0 或 1，则函数返回 1。这处理了基本情况，即没有台阶或只有一阶台阶的情况，此时只有一种方式可以到达楼顶。
		return 1
	}
	d := make([]int, n+1) // 创建了一个整数切片 d，长度为 n+1，并将前两个元素初始化为 1。d 用于存储到达每个台阶级别的方式数。
	d[0] = 1
	d[1] = 1
	for i := 2; i <= n; i++ {
		d[i] = d[i-1] + d[i-2] //    d[i-1] 表示到达第 i-1 层楼梯的方法数。   d[i-2] 表示到达第 i-2 层楼梯的方法数。
	}
	return d[n]
}

func climbStairs(n int) int {
	x, z, r := 0, 0, 1        // x, z, 和 r，分别用于存储计算过程中的中间值   初始化 x 为0，z 为0，r 为1。这代表开始时，到达第0阶楼梯（也就是起点）只有一种方法。
	for i := 1; i <= n; i++ { // 进入循环，从第1阶楼梯开始，一直循环到第 n 阶楼梯
		x = z
		z = r
		r = x + z // 在每次循环中，我们将当前 z 赋给 x，将当前 r 赋给 z，然后将 x 和 z 相加得到新的 r。这表示到达当前阶梯的方法数是前一阶梯和前两阶梯方法数之和
	}
	return r // 最终返回 r，即到达第 n 阶楼梯的方法数。
}
