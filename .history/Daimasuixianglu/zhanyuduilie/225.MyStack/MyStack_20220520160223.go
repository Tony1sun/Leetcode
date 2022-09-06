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

}
