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
		sm[s[r]]
	}

}
