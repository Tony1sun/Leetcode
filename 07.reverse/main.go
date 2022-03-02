package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var num int
	for x != 0 {
		num = num*10 + x%10 // 取余
		x = x / 10          // 取模
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return num
}

func main() {
	x := 123
	fmt.Println(reverse(x))
}
