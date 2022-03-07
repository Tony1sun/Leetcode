package main

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

作者：xiao_ben_zhu
链接：https://leetcode-cn.com/problems/3sum-closest/solution/gu-ding-yi-ge-shu-zai-shuang-zhi-zhen-shun-bian-fu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



func main() {

}