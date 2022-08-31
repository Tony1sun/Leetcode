package main

import (
	"fmt"
	"testing"
)

func Test_Problem144(t *testing.T) {

	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	one.Left = two
	three := &TreeNode{Val: 3}
	two.Left = three

	fmt.Println(preorderTraversal(one))
	fmt.Println(inorderTraversal(one))
}
