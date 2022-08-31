package main

// 225. 用队列实现栈
// https://leetcode.cn/problems/implement-stack-using-queues/
type Mystack struct {
	queue1 []int
	queue2 []int
}

func Constructor() Mystack {
	return Mystack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

func (this *Mystack) Push(x int) {
	// 先将数据存在queue2中
	this.queue2 = append(this.queue2, x)
	// 将queue1所有的元素移到queue2中，再将两个队列交换
	this.Move()
}

func (this *Mystack) Move() {
	if len(this.queue1) == 0 {
		// 交换，queue1置为queue2，queue2置为空
		this.queue1, this.queue2 = this.queue2, this.queue1
	} else {
		// queue1元素从头开始一个一个追加到queue2中
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:] // 去除第一个元素
		this.Move()                   // 重复
	}
}

func (this *Mystack) Pop() {
	val := this.queue1[0]
	this.queue1 = this.queue1[1:]
}
