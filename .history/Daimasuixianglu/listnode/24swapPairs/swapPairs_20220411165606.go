package Leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// 24. 两两交换链表中的节点
// func swapPairs(head *ListNode) *ListNode {
// 	dummy := &ListNode{
// 		Next: head,
// 	}
// 	pre := dummy
// 	for head != nil && head.Next != nil {
// 		pre.Next = head.Next
// 		next := head.Next.Next
// 		head.Next.Next = head
// 		head.Next = next
// 		pre = head
// 		head = next
// 	}
// 	return dummy
// }

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	preNode := dummy
	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		// 交换节点
		first.Next = second.Next
		second.Next = first
		preNode.Next = second

		// 节点往前移动
		preNode = first
		head = first.Next
	}
	return dummy.Next
}
