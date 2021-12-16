package Leetcode

import (
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 3
	node4 := new(ListNode)
	node4.Val = 4
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	head := reverseList(node1)
	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
	fmt.Println()
}
