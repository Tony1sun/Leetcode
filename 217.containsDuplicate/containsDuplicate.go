package Leetcode

import "sort"

// 使用哈希表
// func containsDuplicate(nums []int) bool {
// 	m := make(map[int]bool)
// 	for _, v := range nums {
// 		if _, ok := m[v]; ok {
// 			return true
// 		}
// 		m[v] = true
// 	}
// 	return false
// }

// 排序
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
