package main

import (
	"testing"
)

func Test_Problem144(t *testing.T) {

	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2}
	one.left = two
	three := &TreeNode{Val: 3}
	two.left = three
	four := &TreeNode{Val: 4}
	two.right = four
	five := &TreeNode{v: 5}
	one.right = five
	six := &TreeNode{v: 6}
	five.right = six

	flatten(one)
}
