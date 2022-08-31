package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }


type ListNode = structures

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//声明快慢指针
	slow, fast := head, head
	//链表为nil，直接返回
	if fast == nil {
		return head
	}
	//fast指针先跑n步
	for n > 0 {
		fast = fast.Next
		n--
	}
	//如果这时fast为nil，则head节点是要删除的那个(容易忽略)
	if fast == nil {
		return head.Next
	}
	//快慢指针一起移动
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	//slow.next为需要删除的节点
	slow.Next = slow.Next.Next
	return head
}

func main() {

}
