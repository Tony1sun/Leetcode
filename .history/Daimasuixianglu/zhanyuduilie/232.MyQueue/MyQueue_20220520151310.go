package main

/*
type MyQueue struct {
	stack []int
	back  []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

// 将元素 x 推到队列的末尾
func (this *MyQueue) Push(x int) {
	for len(this.back) != 0 {
		val := this.back[len(this.back)-1]
		this.back = this.back[:len(this.back)-1]
		this.stack = append(this.stack, val)
	}
	this.stack = append(this.stack, x)
}

//  从队列的开头移除并返回元素
func (this *MyQueue) Pop() int {
	for len(this.stack) != 0 {
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back, val)
	}
	if len(this.back) == 0 {
		return 0
	}
	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val
}

// 返回列表开头元素
func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == 0 {
		return 0
	}
	this.back = append(this.back, val)
	return val
}

// 如果队列为空，返回 true ；否则，返回 false
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0 && len(this.back) == 0
}
