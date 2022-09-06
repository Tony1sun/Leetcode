package main

import "strconv"

func evalRPN(tokens []string) int {
	// stack := []int{}
	var stack []int
	for _, token := range tokens {
		val, err := strconv.Atoi(token) // 把字符串转成数字
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}


func main() {
	
}