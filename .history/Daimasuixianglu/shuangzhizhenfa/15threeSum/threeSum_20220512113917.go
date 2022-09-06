package Leetcode

import "sort"

// 三数之和
// https://leetcode-cn.com/problems/3sum/
// https://leetcode-cn.com/problems/3sum/solution/zhi-zhen-yi-dong-guo-cheng-zhong-tiao-guo-zhong-fu/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	// i < len(nums)-2 预留左右指针的位置
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		//若当前i指向元素>0，则代表left和right以及i的和大于0(因为nums已经从小到大排序)。直接break
		if n1 > 0 {
			break
		}
		// 去重
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		// 左右指针
		l, r := i+1, len(nums)-1
		// 当左指针小于右指针时进行循环
		for l < r {
			n2, n3 := nums[l], nums[r]
			// 求出三数之和
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 去重,如果l<r,或者nums[l]==n2，那么l++
				for l < r && nums[l+1] == n2 {
					l++
				}
				for l < r && nums[r-1] == n3 {
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
