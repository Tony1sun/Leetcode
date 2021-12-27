package Leetcode

// 第454题.四数相加II
// 1、首先定义 一个unordered_map，key放a和b两数之和，value 放a和b两数之和出现的次数。
// 2、遍历大A和大B数组，统计两个数组元素之和，和出现的次数，放到map中。
// 3、定义int变量count，用来统计a+b+c+d = 0 出现的次数。
// 4、在遍历大C和大D数组，找到如果 0-(c+d) 在map中出现过的话，就用count把map中key对应的value也就是出现次数统计出来。
// 5、最后返回统计值 count 就可以了
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	count := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += m[0-v3-v4]
		}
	}
	return count
}
