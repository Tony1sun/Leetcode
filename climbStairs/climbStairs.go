package Leetcode

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	pre, cur := 1, 2
	for i := 3; i <= n; i++ {
		j := pre + cur
		pre = cur
		cur = j
	}
	return cur
}
