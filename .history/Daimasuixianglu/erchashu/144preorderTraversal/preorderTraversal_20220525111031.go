package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (res []int) { // 确定递归函数的参数和返回值
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil { // 确定终止条件
			return
		}
		res = append(res, node.Val) 确定单层递归的逻辑
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
