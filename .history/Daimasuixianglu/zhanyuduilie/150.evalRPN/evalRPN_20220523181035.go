package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	// stack := []int{}
	var stack []int
	for _, token := range tokens {
		val, err := strconv.Atoi(token) // 把字符串转成数字
		if err == nil {                 // 如果是数字，入栈
			stack = append(stack, val)
		} else { // 是运算符，取倒数第一、第二个数字来做计算
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2] // 更新stack
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

// func main() {
// 	s := []string{"2", "1", "+", "3", "*"}
// 	fmt.Println(evalRPN(s))
// }

func Test_eva