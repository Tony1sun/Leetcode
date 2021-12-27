package Leetcode

import (
	"fmt"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(nums, 7))
}
