package Leetcode

import "github.com/halfrost/LeetCode-Go/structures"

// https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 02.07. 链表相交
type ListNode = structures.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	// 求A、B长度
	for curA != nil {
		curA = curA.Next
		lenA++
	}

	for curB != nil {
		curB = curB.Next
		lenB++
	}

	var step int
	var fast, slow *ListNode
	// 请求长度差，让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表 遇到相同跳出
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
