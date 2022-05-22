package main

// https://leetcode.cn/problems/valid-parentheses/

func isValid(s string) bool {
	stack := []string{}

	for _, ch := range s {
		c := string(ch)
		// 是左括号，入栈
		if c == "{" || c == "[" || c == "(" {
			stack = append(stack, c)
		} else {
			// 是右括号，栈空，无法匹配
			if len(stack) == 0 {
				return false
			}
			// 获取栈顶
			top := stack[len(stack)-1]
			// 如果栈顶是对应的左括号，被匹配，出栈
			if top == "(" && c == ")" || top == "{" && c == "}" || top == "[" && c == "]" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	// 栈空，则所有的左括号找到匹配；栈中还剩有左括号，则没被匹配
	return len(stack) == 0
}
