package main

// 225. 用队列实现栈
type Mystack struct {
	queue1 []int
	queue2 []int
}


func Constructor() Mystack {
	return Mystack{
		queue1: make([]int,0),
		queu: make([]int,0),
	}
}