package main

import "fmt"

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// 示例 1:
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]

func moveZeroes(nums []int) {
	var l1, l2 []int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			l1 = append(l1, nums[i])
		} else {
			l2 = append(l2, nums[i])
		}
	}
	l2 = append(l2, l1...)
	for i := 0; i < len(nums); i++ {
		nums[i] = l2[i]
	}
}
func main() {
	s := []int{0, 1, 0, 3, 12}
	moveZeroes(s)
	fmt.Println(s)
}
