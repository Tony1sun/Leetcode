package main

import (
	"fmt"
	"testing"
)

func Test_intersect(t *testing.T) {
	nums2 := []int{4, 9, 5}
	nums1 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}
