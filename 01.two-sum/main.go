package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		// 新建当前节点的下一个节点
		curr.Next = new(ListNode)
		// 将curr指向当前节点的下一个节点
		curr = curr.Next
		if l1 != nil {
			carry = carry + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry = carry + l2.Val
			l2 = l2.Next
		}
		// curr.Val = carry % 10
		// carry = carry / 10
		curr.Val, carry = carry%10, carry/10
	}
	return dummy.Next
}

func main() {
	l1 := ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 9}

	l2 := ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	node := addTwoNumbers(&l1, &l2)
	fmt.Printf("%v\n", node)
	fmt.Printf("%v\n", node.Next)
	fmt.Printf("%v\n", node.Next.Next)
}
