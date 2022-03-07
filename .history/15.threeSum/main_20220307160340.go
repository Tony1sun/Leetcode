package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 == 0 {
			break
		}
		// 左右指针
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res,[]int{n1+n2+n3})
			}
		}
	}
}

func main() {
	nums := []int{1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
