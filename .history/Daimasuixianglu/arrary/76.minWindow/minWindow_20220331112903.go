package main

// https://leetcode-cn.com/problems/minimum-window-substring/
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
		if _, ok := sm[s[r]]; ok && sm[s[r]] == tm[s[r]] {
			matchLen++
		}
		// 达到匹配长度时，缩减左右指针直接的长度，得到最短匹配字符串
		for matchLen == len(tm) {
			if len(ans) == 0 || len(s[l:r+1]) < len(ans) {
				ans = s[l : r+1]
			}
			// 左指针指向的字符在t里且数量一致时，匹配长度-1
			if _, ok := sm[s[l]]; ok && sm[s[l]] == tm[s[l]] {
				matchLen--
			}
			// 左指针往前移动
			sm[s[l]]--
			l++
		}
	}
	return ans
}
