package main

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

/*
// 中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left) // 确定单层递归的逻辑,左中右
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
*/
