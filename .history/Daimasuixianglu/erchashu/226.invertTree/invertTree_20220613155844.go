package main

import "github.com/go-playground/locales/root"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历-中左右
func invertTree(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root

}
