package main

// https://leetcode-cn.com/problems/implement-strstr/
// 实现 strStr() 函数。
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。
//
// 说明：
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

// 示例 1：
// 输入：haystack = "hello", needle = "ll"
// 输出：2
//
// 示例 2：
// 输入：haystack = "aaaaa", needle = "bba"
// 输出：-1
//
// 示例 3：
// 输入：haystack = "", needle = ""
// 输出：0

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] { // 不匹配
			j = next[j-1] // 回退到j的前一位,寻找之前匹配的位置
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

func getNext(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] { // 前后缀不相同
			j = next[j-1] // 向前回退
		}
		// 找到相同前后缀
		if s[i] == s[j] {
			j++
		}
		// 将j（前缀的长度）赋给next[i], 因为next[i]要记录相同前后缀的长度
		next[i] = j
	}
}
