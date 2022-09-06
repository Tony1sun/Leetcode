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
	four := &TreeNode{Val: 4}
	two.Right = four
	five := &TreeNode{Val: 5}
	one.Right = five
	six := &TreeNode{Val: 6}
	five.Right = six

	preorderTraversal(one))
}
