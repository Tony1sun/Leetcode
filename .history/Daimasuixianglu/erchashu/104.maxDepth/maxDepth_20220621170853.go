package main

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 104. 二叉树的最大深度
func maxdepth(root *TreeNode) int {
	levl := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		levl++
		// 重新统计队列
		l = len(queue)
	}
	return levl
}
