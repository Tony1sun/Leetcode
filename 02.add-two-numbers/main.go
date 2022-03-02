package main

import "fmt"

// 两数之和
func TwoSum(nums []int, target int) []int {
	map1 := make(map[int]int)
	for index, value := range nums {
		// 如果存在值为target-value的数，将其坐标赋值给index2
		index2, ok := map1[target-value]
		if ok {
			return []int{index2, index}
		}
		// 如果不存在就将该值和下标放入哈希表
		map1[value] = index
	}
	return nil
}

func main() {
	n := []int{1, 2, 7, 11, 15}
	fmt.Println(TwoSum(n, 9))
}
