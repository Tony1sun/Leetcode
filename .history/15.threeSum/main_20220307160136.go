package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 == 0 {
			break
		}
		// 左右指针
		l, r := i+1, len(nums)-1
	}
}

func main() {
	nums := []int{1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
