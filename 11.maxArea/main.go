package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		// 高 取短板的长度
		h := int(math.Min(float64(height[left]), float64(height[right])))
		water := (right - left) * h
		if water > max {
			max = water
		}
		// 向内移动短板，因为若移动长版，则容量可能不变或者变小
		// 所以只能移动短板，容量可能变小也可能变大
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
