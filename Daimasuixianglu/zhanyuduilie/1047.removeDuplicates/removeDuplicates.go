package main

// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
func removeDuplicates(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		// 如果栈不为空
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 入栈
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
