package Leetcode

import "strings"

// 验证回文串
// 先都转成大写（不然会出现 a A 判定为不相同）
// 设置头尾双指针，开启循环
// 如果指向的元素是不是有效的（不是字母和数字），则跳过
// 如果指向的元素有效，但不相同，则不是回文，返回false
// 否则有效，且相同，收缩指针，继续循环
// 直至指针相遇，循环结束，始终没有返回false，返回true。

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if !isValid(s[i]) {
			i++
			continue
		}
		if !isValid(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isValid(c byte) bool {
	op1 := c >= 'a' && c <= 'z'
	op2 := c >= 'A' && c <= 'Z'
	op3 := c >= '0' && c <= '9'
	return op1 || op2 || op3
}
