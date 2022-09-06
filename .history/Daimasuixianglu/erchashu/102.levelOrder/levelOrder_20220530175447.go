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
			node := queue.Remove(queue.Front()).(*TreeNode) // 出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val) // 将值加入本层切片中
		}
		res = append(res, tmpArr) // 放入结果集
		tmpArr = [][]int{}        // 清空层数据
	}
	return res
}
