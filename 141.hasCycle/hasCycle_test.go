package Leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem141(t *testing.T) {

	aa := ListNode{Val: 1}
	bb := ListNode{Val: 2}
	// cc := ListNode{Val: 4}
	// d := ListNode{Val:  -4}
	aa.Next = &bb
	bb.Next = &aa
	// cc.Next = &bb
	//d.Next = &b
	fmt.Println(hasCycle(&aa))
	a := ListNode{Val: 3}
	b := ListNode{Val: 2}
	c := ListNode{Val: 0}
	d := ListNode{Val: -4}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &b
	fmt.Println(hasCycle(&a))
}
