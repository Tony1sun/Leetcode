package main

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

func minDepth(root *TreeNode) int {
	dep := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		dep++
		for ; l > 0; l-- {
			node := queue[0]
			// 如果是叶子节点(左右孩子都为空的节点)
			if node.Left == nil && node.Right == nil {
				return dep
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		l = len(queue)
	}
	return dep
}