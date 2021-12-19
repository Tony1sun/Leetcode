package Leetcode

import "strconv"

// 415. 字符串相加
// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
// 示例 1：
// 输入：num1 = "11", num2 = "123"
// 输出："134"
// 示例 2：

// 输入：num1 = "456", num2 = "77"
// 输出："533"

func addStrings(num1 string, num2 string) string {
	res := ""
	carry := 0
	l1, l2 := len(num1)-1, len(num2)-1
	for l1 >= 0 || l2 >= 0 {
		x, y := 0, 0
		if l1 >= 0 {
			x = int(num1[l1] - byte('0'))
		}
		if l2 >= 0 {
			y = int(num2[l2] - byte('0'))
		}
		sum := x + y + carry
		res += strconv.Itoa(sum % 10)
		carry = sum / 10

		l1--
		l2--
	}
	if carry != 0 {
		res += strconv.Itoa(carry)
	}
	return reverseString(res)
}

// 翻转字符串
func reverseString(str string) string {
	temp := []byte(str)
	left, right := 0, len(temp)-1
	for left < right {
		temp[left], temp[right] = temp[right], temp[left]
		left++
		right--
	}
	return string(temp)
}
