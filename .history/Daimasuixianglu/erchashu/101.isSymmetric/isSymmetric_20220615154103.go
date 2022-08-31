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

// 如果root不等于nil,就把左子树和右子树加入栈
// 循环，左右子树出栈
// 如果左右子树都是nil，continue
// 如果左子树或者右子树或者左子树的Val不等于右子树的Val,