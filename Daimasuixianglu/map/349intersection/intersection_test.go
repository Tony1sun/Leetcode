package Leetcode

import (
	"fmt"
	"testing"
)

func Test_intersection(t *testing.T) {
	nums1 := []int{2, 4, 5, 6, 5, 5}
	nums2 := []int{4, 7, 9, 5}
	fmt.Println(intersection(nums1, nums2))
}
