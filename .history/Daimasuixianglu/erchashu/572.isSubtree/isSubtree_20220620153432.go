package main

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

// https://leetcode.cn/problems/subtree-of-another-tree/

// 判断两个数是否相等
func isEqual(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val == right.Val {
		return isEqual(left.Left, right.Left) && isEqual(right.Right, left.Right)
	}
	return false
}

func isSubtree(left *TreeNode, right *TreeNode) bool {
	// 暴力匹配，从left的根节点开始，依次判定根节点、左节点、右节点的子树是不是等于right
	if left == nil {
		return false
	}
	return isEqual(left, right) || isSubtree(left.Left, right) || isSubtree) 
}
