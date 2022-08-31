package Leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem19(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	node := head.Next
	node.Next = &ListNode{Val: 3}
	node = node.Next
	node.Next = &ListNode{Val: 4}
	node = node.Next
	node.Next = &ListNode{Val: 5}
	d := removeNthFromEnd2(head, 2)
	for d != nil {
		fmt.Println(d.Val)
		d = d.Next
	}
}
