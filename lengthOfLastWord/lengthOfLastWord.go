package Leetcode

// 最后一个单词的长度
// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
// 示例 1：
// 输入：s = "Hello World"
// 输出：5

func lengthOfLastWord(s string) int {
	var ans int
	for i := len(s) - 1; i >= 0; i-- {
		// 当前字符是空格，且之前包含非空格
		if s[i] == ' ' && ans > 0 {
			break
		}
		// 非空格+1
		if s[i] != ' ' {
			ans++
		}
	}
	return ans
}
