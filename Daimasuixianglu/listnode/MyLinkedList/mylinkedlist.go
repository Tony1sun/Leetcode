package main

import "fmt"

// 循环双链表
type MyLinkedList struct {
	dummy *Node
	size  int
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

// 保存哑节点
func Constructor() MyLinkedList {
	rear := &Node{
		Val:  -1,
		Next: nil,
		Pre:  nil,
	}
	rear.Next = rear
	rear.Pre = rear

	return MyLinkedList{rear, 0}
}

// 获取链表长度
func (this *MyLinkedList) GetLength() int {
	return this.size
}

// 是否空链表
func (this *MyLinkedList) IsEmpty() bool {
	return this.size == 0
}

// 获取链表第index个节点的数值
func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.Next
	// head == this，遍历完全
	for head != this.dummy && index > 0 {
		index--
		head = head.Next
	}
	// 否则,head == this,索引无效
	if 0 != index {
		return -1
	}
	return head.Val
}

// 在链表的最前面插入一个节点
func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.dummy
	node := &Node{
		Val: val,
		// head.Next 指向原头节点
		Next: dummy.Next,
		// head.Pre 指向哑节点
		Pre: dummy,
	}
	// 更新原头节点
	dummy.Next.Pre = node
	// 更新哑节点
	dummy.Next = node
	// 更新链表长度
	this.size++
}

// 在链表的最后面插入一个节点
func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummy
	rear := &Node{
		Val: val,
		// 哑节点
		Next: dummy,
		// 指向原尾节点
		Pre: dummy.Pre,
	}
	dummy.Pre.Next = rear
	dummy.Pre = rear
	// 更新链表长度
	this.size++
}

// 在链表第index个节点前面插入一个节点
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	head := this.dummy.Next
	for head != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	node := &Node{
		Val:  val,
		Next: head,
		Pre:  head.Pre,
	}
	head.Pre.Next = node
	head.Pre = node
}

// 删除链表的第index个节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	// 链表为空
	if this.dummy.Next == this.dummy {
		return
	}
	head := this.dummy.Next
	for head.Next != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	// 验证index有效
	if index == 0 {
		head.Next.Pre = head.Pre
		head.Pre.Next = head.Next
	}
	// 更新链表长度
	this.size--
}

// 遍历链表
// func (this *MyLinkedList) traverse() {
// 	var p *Node = this.dummy.Next
// 	if (this.is)

// }

func main() {
	linklist := Constructor()
	fmt.Println("链表创建成功")
	fmt.Println("---开始插入数据---")
	linklist.AddAtHead(1)
	fmt.Println("链表长度为:", linklist.GetLength())
	fmt.Println("---开始插入数据---")
	linklist.AddAtTail(2)
	fmt.Println("链表长度为:", linklist.GetLength())
	// fmt.Println("---开始删除数据---")
	// linklist.DeleteAtIndex(1)
	// fmt.Println("链表长度为:", linklist.GetLength())
	fmt.Println(linklist.Get(0))
}
