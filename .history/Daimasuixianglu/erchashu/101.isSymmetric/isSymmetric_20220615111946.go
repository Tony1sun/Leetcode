package main

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

func isSymmetric(root *TreeNode) bool {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[]
	}
}
