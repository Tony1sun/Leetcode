package Leetcode

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	nums1 := searchInsert(nums, 2)
	fmt.Println(nums1)
}
