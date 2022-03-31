package main

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// 示例 1:
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]

// 方法1. 双指针交换
// 指针 i 、j 初始指向索引 0
// 指针 j 指向提供给非 0 项的坑位
// 指针 i 推进找非 0 项，想交换到 j 指向的坑位

func moveZeroes(nums []int) []int {
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
	return nums

}

// func main() {
// 	nums := []int{0, 1, 0, 3, 12}

// 	fmt.Println(moveZeroes(nums))
// }
