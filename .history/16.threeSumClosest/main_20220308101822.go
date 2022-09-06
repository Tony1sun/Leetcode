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
			// 如果 sum 大于目标值，就右指针左移，使 sum 变小，否则左指针右移，sum 变大。
			if sum > target {
				r--
			} else {
				l++
			}
			再看 abs(sum - target) 是否比之前更小了，如果是，将当前 sum 更新给 res
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
