package Leetcode

import (
	"fmt"
	"testing"
)

func Test_reverseString(t *testing.T) {
	str := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(string(str))
	reverseString(str)
	fmt.Println(string(str))
}
