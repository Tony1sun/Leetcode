package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// εζεΊ
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 == 0 {
			break
		}
	}
}

func main() {
	nums := []int{1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
