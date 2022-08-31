package main

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

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
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
			queue = queue[1:]
		}
		levl++
		l = len(queue)
	}
	return levl
}
