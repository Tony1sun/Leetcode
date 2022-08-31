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
		dep
	}
}
