package main

import (
	"fmt"
	"strings"
)

// 给定一个字符串 s,请你找出其中不含有重复字符的 最长子串的长度。
// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

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
	str1 := "abcabc"
	fmt.Println(lengthOfLongestSubstring(str1))
}
