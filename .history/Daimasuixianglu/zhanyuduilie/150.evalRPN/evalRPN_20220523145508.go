package main

import "strconv"

func evalRPN(tokens []string) int {
	// stack := []int{}
	var stack []int
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]

		}
	}
}
