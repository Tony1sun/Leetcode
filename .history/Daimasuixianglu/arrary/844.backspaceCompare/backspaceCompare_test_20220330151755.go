package main

import (
	"fmt"
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	qs := []question844{

		{
			para844{"ab#c", "ad#c"},
			ans844{true},
		},

		{
			para844{"ab##", "c#d#"},
			ans844{true},
		},

		{
			para844{"a##c", "#a#c"},
			ans844{true},
		},

		{
			para844{"a#c", "b"},
			ans844{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 844------------------------\n")

	for _, q := range qs {
		_, p := q.ans844, q.para844
		fmt.Printf("【input】:%v       【output】:%v\n", p, backspaceCompare(p.s, p.t))
	}
	fmt.Printf("\n\n\n")
}
