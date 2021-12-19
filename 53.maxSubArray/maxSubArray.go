package Leetcode

// 53. 最大子数组和
// 1.贪心算法
// func maxSubArray(nums []int) int {
// 	// 假设最大组合值是第一个
// 	MaxSum := nums[0]
// 	// 移动的组合值
// 	temSum := 0
// 	for i := 0; i < len(nums); i++ {
// 		temSum += nums[i]
// 		if temSum > MaxSum {
// 			MaxSum = temSum
// 		}
// 		// 遇到temSum<0后直接抛弃之前的和(负数加下一个数只会变小,此时局部最优)
// 		if temSum < 0 {
// 			temSum = 0
// 		}
// 	}
// 	return MaxSum
// }

// 2.动态规划
// 动态规划是把一个大问题拆解成一堆小问题
// 该大问题是否能用动态规划解决的是这些小问题会不会被被重复调用
// 如1+1+1+1=4,1+1+1+1+1=4+1=5
// 我们用f(i) 代表以第 i 个数结尾的「连续子数组的最大和」
// 我们只需要求出每个位置的f(i),f(i)=max(f(i-1),nums[i]),然后返回 f 数组中的最大值即可
// 代码中优化了这个数组(直接取max(f,MaxSum))
// 时间复杂度O(n)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSubArray(nums []int) int {
	MaxSum := nums[0]
	f := nums[0]
	for i := 1; i < len(nums); i++ {
		f = max(f+nums[i], nums[i])
		if f > MaxSum {
			MaxSum = f
		}
	}
	return MaxSum
}
