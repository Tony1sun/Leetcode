package Leetcode

// 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 注意：给定 n 是一个正整数。

// 示例 1：
// 输入： 2
// 输出： 2
// 解释： 有两种方法可以爬到楼顶。
// 1.  1 阶 + 1 阶
// 2.  2 阶

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	pre, cur := 1, 2
	for i := 3; i <= n; i++ {
		// 数列从第3项开始，每一项都等于前两项之和。
		j := pre + cur
		// pre和cur往后移
		pre = cur
		cur = j
	}
	return cur
}
