package Leetcode

import "github.com/halfrost/LeetCode-Go/structures"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	var pre *ListNode
	cur := result

	i := 1
	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return result.Next

}


func removeNthFromEnd(head *ListNode, n int) *ListNode {
    newHead := &ListNode{Next: head}
    slow, fast := newHead, newHead
    i := 0
    for fast.Next != nil {
        if i < n {
            fast = fast.Next
            i++
        } else {
            fast = fast.Next
            slow = slow.Next
        }
    }
    slow.Next = slow.Next.Next
    return newHead.Next
}

作者：dfzhou6
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/solution/jian-dan-yi-dong-go-by-dfzhou6-3nh7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。