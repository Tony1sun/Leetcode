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
		
	}
}

func main() {
	nums := []int{1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
