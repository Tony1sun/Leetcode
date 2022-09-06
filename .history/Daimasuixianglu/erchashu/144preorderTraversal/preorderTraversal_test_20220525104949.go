package main

import (
	"testing"
)

func Test_Problem144(t *testing.T) {

	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	one.Left = two
	three := &TreeNode{Val: 3}
	two.Left = three
	four := &TreeNode{Val: 4}
	two.R = four
	five := &TreeNode{Val: 5}
	one.right = five
	six := &TreeNode{Val: 6}
	five.right = six

	flatten(one)
}
