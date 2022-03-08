package Leetcode

import "github.com/halfrost/LeetCode-Go/structures"

。


/*
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	var pre *ListNode
	cur := result

	i := 1
	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return result.Next

}
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{Next: head}
	slow, fast := newHead, newHead
	i := 0
	for fast.Next != nil {
		if i < n {
			fast = fast.Next
			i++
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}
