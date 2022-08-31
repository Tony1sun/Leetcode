package main

import "strconv"

func evalRPN(tokens []string) int {
	// stack := []int{}
	var stack []int
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		}
	}
}
