package Leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem24(t *testing.T) {

	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	d := &ListNode{Val: 4}
	a.Next = b
	b.Next = c
	c.Next = d
	h := swapPairs(a)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}
