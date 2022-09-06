package main

import "github.com/halfrost/LeetCode-Go/structures"

// https://leetcode.cn/problems/symmetric-tree/

type TreeNode = structures.TreeNode

// 迭代法
func isSymmetric(root *TreeNode) bool {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}

// 递归法
func isSymmetric1(root *TreeNode) bool {
	return defs(root.Left, root.Right)
}
