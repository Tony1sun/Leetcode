package Leetcode

// 1. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，
// 并返回他们的数组下标。
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, value := range nums {
		targetNum := target - value
		index2, ok := m[targetNum]
		if ok {
			return []int{index2, index}
		}
		m[value] = index
	}
	return nil
}
