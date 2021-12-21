package Leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next // 将 prev 的 next 改为 head 的 next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		pre = head
		head = next
	}
	return dummy
}
