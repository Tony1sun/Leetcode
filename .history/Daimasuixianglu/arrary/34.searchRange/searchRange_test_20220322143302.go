package main

import (
	"fmt"
	"testing"
)

func Test_searchRange(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchRange(nums, 7))
}
