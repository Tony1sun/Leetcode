package main

import "fmt"

func longestPalindrome(s string) string {
	// 双指针
	start, end := 0, 0
	for i := 0; i < len(s); {
		l, r := i, i

		// 回文子串是偶数或者有连续相同的情况
		for r < len(s)-1 && s[r] == s[r+1] {
			r++
		}
		i = r + 1
		// 两边一起扩展
		for l > 0 && r < len(s)-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		// 更新最长回文子串的长度和起始点
		if end-start < r-l {
			start = l
			end = r
		}
	}
	return s[start : end+1]
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

