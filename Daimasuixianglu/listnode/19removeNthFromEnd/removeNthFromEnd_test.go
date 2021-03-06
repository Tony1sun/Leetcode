package Leetcode

import (
	"fmt"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	d := &ListNode{Val: 4}
	e := &ListNode{Val: 5}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	h := removeNthFromEnd(a, 2)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}
