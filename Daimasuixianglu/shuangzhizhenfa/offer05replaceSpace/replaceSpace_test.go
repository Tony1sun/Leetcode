package Leetcode

import (
	"fmt"
	"testing"
)

type question541 struct {
	para541
	ans541
}

// para 是参数
// one 代表第一个参数
type para541 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans541 struct {
	one string
}

func Test_Problem541(t *testing.T) {

	qs := []question541{
		{
			para541{"We are happy."},
			ans541{"We%20are%20happy."},
		},
		{
			para541{"     "},
			ans541{"%20%20%20%20%20"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 541------------------------\n")

	for _, q := range qs {
		_, p := q.ans541, q.para541
		fmt.Printf("【input】:%v       【output】:%v\n", p, replaceSpace(p.s))
	}
	fmt.Printf("\n\n\n")
}
