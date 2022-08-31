package main

import (
	"container/list"

	"github.com/halfrost/LeetCode-Go/structures"
)

// https://leetcode.cn/problems/binary-tree-level-order-traversal/

type TreeNode = structures.TreeNode

func levelOrder(root *TreeNode) [][]int {
	// 结果集
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)

	var tmpArr []int
	
}
