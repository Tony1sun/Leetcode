package Leetcode

import (
	"fmt"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	n := climbStairs(4)
	fmt.Printf("有%d种方法爬到楼顶", n)
}
