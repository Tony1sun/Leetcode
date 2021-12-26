package Leetcode

import (
	"fmt"
	"testing"
)

func Test_removeElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
}
