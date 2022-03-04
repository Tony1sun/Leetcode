package Leetcode

import (
	"fmt"
	"strings"
)

// 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。
// 示例 1：
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：
// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0]
	for _, value := range strs {
		// 如果result在value出现 子串result在字符串value中第一次出现的位置
		for strings.Index(value, result) != 0 {
			// 消减result，直到result和value匹配
			result = result[0 : len(result)-1]
			if len(result) == 0 {
				return ""
			}
		}
	}
	return result
}

func main() {
	strs := []string{"flower", "flowera", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
