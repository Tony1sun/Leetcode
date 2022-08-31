package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
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

	}
}

// 后序遍历
