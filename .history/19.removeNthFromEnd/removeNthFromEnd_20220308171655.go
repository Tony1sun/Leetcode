package Leetcode

import "github.com/halfrost/LeetCode-Go/structures"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head

	var pre *ListNode
	cur := result

	i := 1
	for head != nil {
		if i >= n {}
	}
}
