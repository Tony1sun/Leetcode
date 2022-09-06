package Leetcode

// 搜索插入位置
// 暴力破解
// func searchInsert(nums []int, target int) int {
// 	for i := 0; i < len(nums); i++ {
// 		if target <= nums[i] {
// 			return i
// 		}
// 	}
// 	return len(nums)
// }

// 借助游标法
func searchInsert(nums []int, target int) int {
	var left = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < target {
			left++
		} else {
			return left
		}
	}
	return left + 1
}

// 二分法
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		middle := (right + left) / 2
		if nums[middle] >= target {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}
