package main

// https://leetcode-cn.com/problems/backspace-string-compare/solution/844zhan-mo-ni-yu-kong-jian-geng-you-de-shuang-zhi-/

// 双指针

// func getString(s string) string {
// 	bz := []rune{}
// 	for _, c := range s {
// 		if c != '#' {
// 			bz = append(bz, c)
// 		} else if len(bz) > 0 {
// 			bz = bz[:len(bz)-1]
// 		}
// 	}
// 	return string(bz)
// }

// func backspaceCompare(s string, t string) bool {
// 	return getString(s) == getString(t)
// }

func backspaceCompare(s string, t string) bool {
	skips, skipt := 0, 0
}
