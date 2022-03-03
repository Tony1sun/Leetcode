package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intSlice := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	s := ""

	for i := len(intSlice) - 1; i >= 0; i-- {
		// 如果num小于当前索引,继续往前移动
		if num < intSlice[i] {
			continue
		}
		// num / 当前索引的数
		repeats := num / intSlice[i]
		// 返回count个roman[i]串联的字符串。
		s += strings.Repeat(roman[i], repeats)
		// num 减去 repeats
		num -= repeats * intSlice[i]
		if num == 0 {
			break
		}
	}
	return s
}

func main() {
	n := 999
	fmt.Println(intToRoman(n))
}
