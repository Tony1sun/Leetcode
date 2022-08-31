package main

// https://leetcode-cn.com/problems/design-linked-list/
// 单链表
type MyLinkedList struct {
	head *Node
	tail *Node
	len  int
}

type Node struct {
	val  int
	next *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// 若index不合法（index < 0 || index >= list.len）,
// 返回-1；否则从头往后找到第index个节点，返回其值。
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.len {
		return -1
	}
	p := this.head
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p.val
}

// 新增节点的next指向现在的头，新的头指向新增的节点；
// 如果是空链表第一次新增的情形，须头尾均指向新增节点。最后将链表长度加1。
func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{val, nil}
	if 0 == this.len {
		this.head = node
		this.tail = node
	} else {
		node.next = this.head
		this.head = node
		node.next, thi
	}
	this.len++
}

// 现在的尾的next指向新增节点，新的尾指向新增的节点；
// 同样须注意在空链表第一次新增的情形，须头尾均指向新增节点。最后将链表长度加1。
func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{val, nil}
	if 0 == this.len {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.len++
}

// addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val 的节点。
// 如果 index 等于链表的长度，则该节点将附加到链表的末尾。
// 如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.len {
		return
	}
	node := &Node{val, nil}
	// 如果index在链表头
	if 0 == index {
		node.next = this.head
		this.head = node
		if 0 == this.len {
			this.tail = node
		}
		// 如果 index 在链表尾部
	} else if this.len == index {
		this.tail.next = node
		this.tail = node
	} else {
		// 若是在链表中间某处新增，则找到新增位置的前一个节点，
		// 将新增的节点插入并建立新的连接关系。最后将链表长度加1
		p := this.head
		for i := 0; i < index-1; i++ {
			p = p.next
		}
		node.next = p.next
		p.next = node
	}
	this.len++
}

// deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.len {
		return
	}
	// 首先判断是否删除的是第一个节点，
	// 若是，则将头指向当前头的下一个节点（若链表仅有一个节点，须同时将尾也指向空）
	if 0 == index {
		this.head = this.head.next
		if 1 == this.len {
			this.tail = nil
		}
	} else {
		p := this.head
		// 找到待删除节点的前一个节点
		for i := 0; i < index-1; i++ {
			p = p.next
		}
		// 将待删节点删除并建立新的连接关系（若删除的是最后一个节点，须同时将尾指向待删除节点的前一个节点）。最后将链表长度减1。
		p.next = p.next.next
		if this.len-1 == index {
			this.tail = p
		}
	}
	this.len--
}
