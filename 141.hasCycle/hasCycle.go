package Leetcode

import "github.com/halfrost/LeetCode-Go/structures"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

type ListNode = structures.ListNode

// 快慢指针
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil { // 快指针指向真实节点
		if fast.Next == nil { // 说明没有环
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next // 慢的走一步，快的走两步
		if slow == fast {     // 快慢指针相遇，有环
			return true
		}
	}
	return false // fash走出去了，没有环
}
