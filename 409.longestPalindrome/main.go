package main

import "fmt"

// 409. 最长回文串
// 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
// 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
// 注意:
// 假设字符串的长度不会超过 1010。
// 示例 1:
// 输入:
// "abccccdd"
// 输出:
// 7
// 解释:
// 我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

func longestPalindrome(s string) int {
	// 初始用于计数的数组，因为只有大小写英文字符
	d := make([]int, 256) // 所以把数组设置成ASCII字符集的大小
	oddNum := 0           // 定义数量是奇数的个数
	// 遍历字符串，统计每个字符出现的次数
	for _, c := range s {
		d[c]++
	}
	// 接下来统计字符出现的次数
	for _, count := range d {
		if count%2 == 1 {
			oddNum++
		}
	}
	// 不能使用的字符，在0和odd中取较大值
	unUsed := 0
	if oddNum > 0 {
		unUsed = oddNum - 1
	}
	// 最后用字符串的长度减去无用字符串的个数
	nums := len(s) - unUsed
	return nums
}

func main() {
	str := "abuyccccdd"
	str1 := longestPalindrome(str)
	fmt.Println(str1)
}
