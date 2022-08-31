package main

import (
	"container/list"

	"github.com/halfrost/LeetCode-Go/structures"
)

type TreeNode = structures.TreeNode

// 递归法
/*
// 中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left) // 确定单层递归的逻辑,左中右
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
*/

// 迭代法-左中右
func inorderTraversal(root *TreeNode) (res []int) {
	var ans []int
	if root == nil {
		return ans
	}
	st := list.New()
	cur := root
	for cur != nil || st.Len() > 0 {
		if cur != nil {
			st.PushBack(cur) // 将一个值为v的新元素插入链表的最后一个位置，返回生成的新元素。
			cur = cur.Left
		} else {
			// 删除链表中的元素，并返回Value。Back 返回链表最后一个元素或nil
			cur = st.Remove(st.Back()).(*TreeNode)
			ans = append(ans, cur.Val)
			cur = cur.Right
			cur = cur.Right
			cur = cur.Right
			cur = cur.Right
		}
	}
	return ans
}
