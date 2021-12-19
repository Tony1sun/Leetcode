package Leetcode

import (
	"fmt"
	"testing"
)

func Test_removeElement(t *testing.T) {
	nums := []int{-1, 0, 3, 9, 9, 12, 14}
	target := 9
	fmt.Println(removeElement(nums, target))
}
