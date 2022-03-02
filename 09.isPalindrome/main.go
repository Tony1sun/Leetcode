package main

import "fmt"

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 例如，121 是回文，而 123 不是。
// 示例 1：
// 输入：x = 121
// 输出：true

func isPalindrome(x int) bool {
	// 倒序后，判断是不是和原来的数字相等
	if x < 0 {
		return false
	}
	origin := x
	redirect := 0
	for x != 0 {
		redirect = redirect*10 + x%10
		x = x / 10
	}
	return redirect == origin
}

func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}
