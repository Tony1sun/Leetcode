package Leetcode

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search1(nums, target))
}
