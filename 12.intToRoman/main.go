package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	roman := []string{"I", "IV", "V", "IV", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intSlice := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	s := ""

	for i := len(intSlice) - 1; i >= 0; i-- {
		if num < intSlice[i] {
			continue
		}
		repeats := num / intSlice[i]
		s += strings.Repeat(roman[i], repeats)
		num -= repeats * intSlice[i]
		if num == 0 {
			break
		}
	}
	return s
}

func main() {
	n := 3
	fmt.Println(intToRoman(n))
}
