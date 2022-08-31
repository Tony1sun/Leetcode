package Leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 142.环形链表II
// 判断链表是否环
// 如果有环，如何找到这个环的入口
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 如果有环
		if slow == fast {
			// 从头结点出发一个指针，从相遇节点 也出发一个指针，这两个指针每次只走一个节点
			// 那么当这两个指针相遇的时候就是 环形入口的节点。
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}
