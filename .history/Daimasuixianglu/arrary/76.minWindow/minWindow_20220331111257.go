package main

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
		// da'dao
	}

}
