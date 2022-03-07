package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[len(nums)-1]

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			sum := n1 + n2 + n3
			if sum > target {
				r--
			} else {
				l++
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
		}
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 1))
}
