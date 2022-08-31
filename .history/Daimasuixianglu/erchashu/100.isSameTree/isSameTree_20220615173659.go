package main

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode


https://leetcode.cn/problems/same-tree/

func isSameTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSameTree(left.Left, right.Left) && isSameTree(right.Right, left.Right)
}
