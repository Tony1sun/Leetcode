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

}
