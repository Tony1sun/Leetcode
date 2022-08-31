package main

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

/*
// 后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left) // 确定单层递归的逻辑,左右中
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}
*/

func postorderTraversal(root *TreeNode) (res []int) {
	var ans []int
	if root == nil {
		return ans
	}
}
