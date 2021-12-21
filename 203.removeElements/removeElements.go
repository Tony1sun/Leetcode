package Leetcode

import "github.com/halfrost/LeetCode-Go/structures"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

type ListNode = structures.ListNode

// 迭代
// 题目要求删除链表中给定值的所有节点，一般的步骤无非就是：
// 遍历链表找到所有值为 val 的节点；
// 值为 val 的节点的上一个节点直接指向该节点的下一个节点（ pre->next = pre->next->next）。

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	pre := head
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return head
}

// 双指针
//设置两个指针分别指向头节点，pre （记录待删除节点的前一节点）和 cur (记录当前节点)；
//遍历整个链表，查找节点值为 val 的节点，找到了就删除该节点，否则继续查找。
//找到，将当前节点的前一节点（之前最近一个值不等于 val 的节点）连接到当前节点（cur 节点）的下一个节点（pre->next = cur->next）。
//没找到，更新最近一个值不等于 val 的节点（pre = cur），继续遍历（cur = cur->next）。
// func removeElements(head *ListNode, val int) *ListNode {
// 	for head != nil {
// 		if head.Val != val {
// 			break
// 		}
// 		head = head.Next
// 	}
// 	pre, cur := head, head
// 	for cur != nil {
// 		if cur.Val == val {
// 			pre.Next = cur.Next
// 		} else {
// 			pre = cur
// 		}
// 		cur = cur.Next
// 	}
// 	return head
// }
