package main

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	j := 0
	next := make([]int, n)
	next[0] = j
	for i := 0; i < n; i++ {
		for j > 0 && s[i] != s
	}
}
