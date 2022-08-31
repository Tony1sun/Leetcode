package main

// https://leetcode-cn.com/problems/backspace-string-compare/solution/844zhan-mo-ni-yu-kong-jian-geng-you-de-shuang-zhi-/

func getring(s string) string {
	bz := []rune{}
	for _, c := range s {
		if c != '#' {
			bz = append(bz, c)
		} else if len(bz) > 0 {
			bz = bz[:len(bz)-1]
		}
	}
	return string(bz)
}
