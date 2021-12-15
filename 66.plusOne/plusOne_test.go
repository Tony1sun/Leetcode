package Leetcode

import (
	"fmt"
	"testing"
)

func Test_plusOne(t *testing.T) {
	slice := []int{0, 1, 9, 9}
	slice1 := plusOne(slice)
	fmt.Println(slice1)
}
