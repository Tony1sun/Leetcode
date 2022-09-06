package main

import (
	"testing"
)

type question144 struct {
	para144
	ans144
}

// para 是参数
// one 代表第一个参数
type para144 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans144 struct {
	one []int
}

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
