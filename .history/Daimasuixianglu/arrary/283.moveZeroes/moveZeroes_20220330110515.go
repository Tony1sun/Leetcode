package main



func moveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			i++
		}
	}
}
