package Leetcode

import (
	"fmt"
	"testing"
)

func Test_removeElements(t *testing.T) {
	l1 := ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	node := removeElements(&l1, 4)
	fmt.Printf("%v\n", node)
	fmt.Printf("%v\n", node.Next)
	fmt.Printf("%v\n", node.Next.Next)
}
