package Leetcode

// 349. 两个数组的交集
func intersection(nums1, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] = 1
	}
	var res []int
	// 利用count>0，实现重复值只拿一次放入返回结果中
	for _, v := range nums2 {
		if count, ok := m[v]; ok && count > 0 {
			res = append(res, v)
			m[v]--
		}
	}
	return res
}

//优化版，利用set，减少count统计
// func intersection1(nums1, nums2 []int) []int {
// 	set := make(map[int]struct{}, 0)
// 	res := make([]int, 0)
// 	for _, v := range nums1 {
// 		if _, ok := set[v]; !ok {
// 			set[v] = struct{}{}
// 		}
// 	}
// 	for _, v := range nums2 {
// 		if _, ok := set[v]; ok {
// 			res = append(res, v)
// 			delete(set, v)
// 		}
// 	}
// 	return res
// }
