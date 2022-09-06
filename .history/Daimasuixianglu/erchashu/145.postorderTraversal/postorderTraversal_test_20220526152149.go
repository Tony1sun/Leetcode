package main

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

func Test_Problem144(t *testing.T) {

	qs := []question144{

		{
			para144{[]int{}},
			ans144{[]int{}},
		},

		{
			para144{[]int{1}},
			ans144{[]int{1}},
		},

		{
			para144{[]int{1, structures.NULL, 2, 3}},
			ans144{[]int{1, 2, 3}},
		},
		{
			para144{[]int{1, 2, 3, 4, 5, 6}},
			ans144{[]int{1, 2, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 144------------------------\n")

	for _, q := range qs {
		_, p := q.ans144, q.para144
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", preorderTraversal(root))
	}
	fmt.Printf("\n\n\n")
}
