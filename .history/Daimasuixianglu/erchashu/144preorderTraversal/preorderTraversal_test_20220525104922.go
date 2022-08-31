package main

import (
	"testing"
)



func Test_Problem144(t *testing.T) {

	one := &TreeNode{v: 1}
	two := &TreeNode{v: 2}
	one.left = two
	three := &TreeNode{v: 3}
	two.left = three
	four := &TreeNode{v: 4}
	two.right = four
	five := &TreeNode{v: 5}
	one.right = five
	six := &TreeNode{v: 6}
	five.right = six

	flatten(one)
}
