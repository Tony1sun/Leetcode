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

func (this *MyLinkedList) Get(index int) int {
	head := this.head
	for i := 0; i < index && head != nil; i++ {
		head = head.Next
	}
	if head != nil {
		
	}
}
