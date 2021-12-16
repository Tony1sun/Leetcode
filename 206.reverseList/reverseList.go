package Leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		// 保存头节点的下一个节点
		next := cur.Next
		// 将头节点指向前一个节点
		cur.Next = pre
		// 更新前一个节点
		pre = cur
		// 更新头节点
		cur = next
	}
	return pre
}
