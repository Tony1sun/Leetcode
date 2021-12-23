package Leetcode

// 242.有效的字母异位词
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 示例 1: 输入: s = "anagram", t = "nagaram" 输出: true
// 示例 2: 输入: s = "rat", t = "car" 输出: false
// 说明: 你可以假设字符串只包含小写字母。

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	exists := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if v, ok := exists[s[i]]; v >= 0 && ok {
			//
			exists[s[i]] = v + 1
		} else {
			exists[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		if v, ok := exists[t[i]]; v >= 1 && ok {
			exists[t[i]] = v - 1
		} else {
			return false
		}
	}
	return true
}
