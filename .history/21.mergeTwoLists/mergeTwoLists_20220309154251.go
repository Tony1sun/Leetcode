package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for l1 != nil || l2 != nil {
		if l1 == nil {
			tail.Next = l2
			return head.Next
		}

		if l2 == nil {
			tail.Next = l1
			return head.Next
		}

		if l1.Val > l2.Val {
			tail.Next = l2
			l2 = l2.Next
			tail = tail.Next
			tail.Next = nil
		} else {
			tail.Next = l1
			l1 = l1.Next
			tail = tail.Next
			tail.Next = nil
		}
	}
	return nil
}
*/

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = &ListNode{2, &ListNode{4, nil}}

	l2 := new(ListNode)
	l2.Val = 1
	l2.Next = &ListNode{3, &ListNode{4, nil}}

	m := mergeTwoLists(l1, l2)

	for {
		fmt.Print(m.Val)
		if m.Next == nil {
			break
		}
		m = m.Next
	}
}
