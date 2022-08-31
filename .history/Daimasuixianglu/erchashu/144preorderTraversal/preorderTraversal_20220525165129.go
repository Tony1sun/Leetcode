package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
// 前序遍历
/*
func preorderTraversal(root *TreeNode) (res []int) { // 确定递归函数的参数和返回值
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil { // 确定终止条件
			return
		}
		res = append(res, node.Val) // 确定单层递归的逻辑,中左右
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

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

// 后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left) // 确定单层递归的逻辑,左右中
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}
*/

// 迭代法
// 前序遍历-中左右
func preorderTraversal(root *TreeNode) (res []int) {
	ans := []int{0}
	if root == nil {
		return ans
	}
	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Right !
	}
}
