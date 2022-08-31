package Leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 19.删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	cur := head
	prev := dummyHead.Next
	i := 1
	// cur 先走n+1步
	for cur != nil {
		cur = cur.Next
		if i > n {
			// 往前移动一位
			prev = prev.Next
		}
		i++
	}
	// cur 走到结尾后，prev的下一个结点为倒数第N个结点
	prev.Next = prev.Next.Next
	return dummyHead.Next
}
