package Leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/reverse-linked-list/
// 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		// 获取下一个结点
		next := cur.Next
		// 将下一个结点指向前一个节点
		cur.Next = pre
		// 将当前结点赋值给前结点
		pre = cur
		// 将next结点赋值给当前结点，用于继续往下执行
		cur = next
	}
	return pre
}
