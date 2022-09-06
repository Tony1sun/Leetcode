package main

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/subtree-of-another-tree/

func isSameTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isSameTree(left.Left, right.Left) && isSameTree(right.Right, left.Right)
}
