package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	// 记录字串长度
	var length int
	left, right := 0, 0
	s1 := s[left:right]
	for right := 0; right < len(s); right++ {
		// 记录right在s1中出现的位置
		index := strings.IndexByte(s1, s[right])
		if index != -1 {
			left += index + 1
		}
		// 窗口右移
		s1 = s[left : right+1]
		if len(s1) > length {
			length = len(s1)
		}
	}
	return length
}

func main() {
	str := "abcabcdd"
	fmt.Println(lengthOfLongestSubstring(str))
}
