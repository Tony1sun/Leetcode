package main

type MyLinkedList struct {
	head *Node
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
func (this *MyLinkedList) Get(index int) int {
	head := this.head
	for i := 0; i < index && head != nil; i++ {
		head = head.Next
	}
	if head != nil {
		return head.Val
	} else {
		return -1
	}
}

// addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
}

// addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int)