package Leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

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

// [1>2>3>4]
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	preNode := dummy
	for head != nil && head.Next != nil {
		first := head
		second := head.Next
		// 交换前两个节点
		first.Next = second.Next
		second.Next = first
		preNode.Next = second
		// 把位置移到前一个节点去
		preNode = first
		head = first.Next
	}
	return dummy.Next
}
