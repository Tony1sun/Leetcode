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
	for queue.Len() > 0 {
		length := queue.Len() // 保存当前层的长度， 然后处理当前层(十分重要，防止添加下层元素影响判断层中元素的个数)
		for i := 0; i < length; i++ {
			
		}
	}
}
