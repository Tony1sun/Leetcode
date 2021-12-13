package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 头节点
	dummy := new(ListNode)
	// 指向头节点的临时节点
	cur := dummy
	// 进位
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		// 新建头节点的下一个节点
		cur.Next = new(ListNode)
		// 将cur指向当前节点的下一个节点
		cur = cur.Next
		if l1 != nil {
			carry = carry + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry = carry + l2.Val
			l2 = l2.Next
		}
		// cur.Val = carry % 10
		// carry = carry / 10
		// %10是求余, /10是求倍
		cur.Val, carry = carry%10, carry/10
	}
	return dummy.Next
}

func main() {
	l1 := ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	node := addTwoNumbers(&l1, &l2)
	fmt.Printf("%v\n", node)
	fmt.Printf("%v\n", node.Next)
	fmt.Printf("%v\n", node.Next.Next)
}
