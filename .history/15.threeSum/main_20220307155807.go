package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// εζεΊ
	sort.Ints(nums)

	
}

func main() {
	nums := []int{1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
