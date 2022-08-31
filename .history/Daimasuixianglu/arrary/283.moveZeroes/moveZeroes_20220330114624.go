package main

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// 示例 1:
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]

func moveZeroes(nums []int) []int {
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 { // i遇到非0
			nums[i], nums[j] = nums[j], nums[i] // 和j交换，安排到j的位置
			i++                                 // i推进，找下一个非0
			j++                                 // 安排了一个，j后移
		} else { // i遇到0，继续后移，找非0。j 继续指向待安排的坑位
			i++
		}
	}
	return nums
}

// func main() {
// 	nums := []int{0, 1, 0, 3, 12}

// 	fmt.Println(moveZeroes(nums))
// }
