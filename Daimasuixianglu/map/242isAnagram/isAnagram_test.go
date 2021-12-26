package Leetcode

import (
	"fmt"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	s := "anagram"
	y := "nagaram"
	fmt.Println(isAnagram(s, y))
}
