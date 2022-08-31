package main

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

func defs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	// if left.Val != right.Val {
	// 	return false
	// }
	return left.Val != right.Val && defs(left.Left, right.Left) && defs(right.Right, )
}
func isSameTree(root *TreeNode) bool {
	return defs(root.Left, root.Right)
}
