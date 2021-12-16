package Leetcode

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 1}
	l1.Next.Next = &ListNode{Val: 2}
	l1.Next.Next.Next = &ListNode{Val: 3}

	node := deleteDuplicates(&l1)
	fmt.Printf("%v\n", node)
	fmt.Printf("%v\n", node.Next)
	fmt.Printf("%v\n", node.Next.Next)
}
