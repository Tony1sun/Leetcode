package main

import "fmt"

func longestPalindrome(s string) string {
	//双指针法
	start, end := 0, 0 //最长回文子串的起始\末尾位置

	for i := 0; i < len(s);  {
		// 左右指针，向两边扩展
		l, r := i, i

		// 回文子串可能是偶数的情况，让r先走
		for r < len(s)-1 && s[r] == s[r+1] { //防止越界，一直向后比较，不相等停止
			r++
		}
		// i到达r所扩展的最远长度的下一个字符w
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
	str := "adac"
	str1 := longestPalindrome(str)
	fmt.Println(str1)
}
