package Leetcode

// 二分查找
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target，
// 写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
// 示例：输入: nums = [-1,0,3,5,9,12], target = 9
// 输出: 4
// 解释: 9 出现在 nums 中并且下标为 4

func search(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func search1(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	for low <= high {
		mid := (high+low) / 2
		if nums[mid] > target {

		} else if nums[mid] < target {
			high = mid - 1
		} else {

		}
	}
}
