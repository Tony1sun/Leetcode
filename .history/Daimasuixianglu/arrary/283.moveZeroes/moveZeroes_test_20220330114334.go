package main

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	n := []int{0, 1, 0, 3, 12}
	
	fmt.Println(moveZeroes(n))
}
