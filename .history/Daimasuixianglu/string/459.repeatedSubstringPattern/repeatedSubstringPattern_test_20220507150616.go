package main

import (
	"fmt"
	"testing"
)

type question28 struct {
	para28
	ans28
}

// para 是参数
// one 代表第一个参数
type para28 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans28 struct {
	one bool
}

func Test_Problem28(t *testing.T) {

	qs := []question28{
		{
			para28{"aabaabaafa"},
			ans28{true},
		},

		{
			para28{"abab"},
			ans28{true},
		},

		{
			para28{"aba"},
			ans28{},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 28------------------------\n")

	for _, q := range qs {
		_, p := q.ans28, q.para28
		fmt.Printf("【input】:%v       【output】:%v\n", p, repeatedSubstringPattern(p.s))
	}
	fmt.Printf("\n\n\n")
}
