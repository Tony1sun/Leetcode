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

作者：xiao_ben_zhu
链接：https://leetcode-cn.com/problems/3sum/solution/zhi-zhen-yi-dong-guo-cheng-zhong-tiao-guo-zhong-fu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	nums := []int{1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
