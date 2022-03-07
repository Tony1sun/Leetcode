package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		排序后，如果外层遍历的数已经大于0，则另外两个数一定大于0，sum不会等于0，直接break
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 找到一个解后，左右指针同时向内收缩，为了避免指向重复的元素，需要：
				// 左指针在保证left < right的前提下，一直右移，直到指向不重复的元素
				// 右指针在保证left < right的前提下，一直左移，直到指向不重复的元素
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
