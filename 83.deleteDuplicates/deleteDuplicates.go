package Leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除排序链表中的重复元素
// 递归
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	head.Next = deleteDuplicates(head.Next)
// 	// 如果头结点和下一个结点相等
// 	if head.Val == head.Next.Val {
// 		head.Next = head.Next.Next
// 	}
// 	return head
// }

// 双指针
// 从头开始遍历，记录当前结点的值temp，如果下一个结点的值
// 与当前的值相同就把下一个节点删除，否则就移动当前结点为
// 下一个结点，更新当前值为下一个结点的值。
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	temp := cur.Val

	for cur.Next != nil {
		if temp != cur.Next.Val {
			cur = cur.Next
			temp = cur.Val
		} else {
			cur.Next = cur.Next.Next
		}
	}
	return head
}
