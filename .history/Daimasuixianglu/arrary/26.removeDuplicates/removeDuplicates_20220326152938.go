package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	low := 0
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[low] {
			nums[low+1] = nums[fast]
			low++
		}
	}
	return low+=
}
