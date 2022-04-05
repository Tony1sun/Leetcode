package main

// https://leetcode-cn.com/problems/minimum-window-substring/
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

func minWindow(s, t string) string {
	tm, sm := map[byte]int{}, map[byte]int{}
	var matchLen int
	var ans string

	// 计算t的字符
	for i := range t {
		tm[t[i]]++
	}
	// 依次遍历s,l为左指针,r为右指针
	for l, r := 0, 0; r < len(s); r++ {
		// 右指针指向的字符在t里且数量一致时，匹配长度+1
		sm[s[r]]++
		if _, ok := tm[s[r]]; ok && sm[s[r]] == tm[s[r]] {
			matchLen++
		}
		// 达到匹配长度时，缩减左右指针之间的长度，得到最短匹配字符串
		for matchLen == len(tm) {
			if len(ans) == 0 || len(s[l:r+1]) < len(ans) {
				ans = s[l : r+1]
			}
			// 左指针指向的字符在t里且数量一致时，匹配长度-1
			if _, ok := tm[s[l]]; ok && sm[s[l]] == tm[s[l]] {
				matchLen--
			}
			// 左指针往前移动
			sm[s[l]]--
			l++
		}
	}
	return ans
}
