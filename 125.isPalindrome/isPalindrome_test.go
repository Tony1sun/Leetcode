package Leetcode

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	str := "race a car"
	isTure := isPalindrome(str)
	fmt.Println("当前字符串是否是回文串？", isTure)
}
