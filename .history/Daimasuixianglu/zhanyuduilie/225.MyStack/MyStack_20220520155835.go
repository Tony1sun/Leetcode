package main

// 225. 用队列实现栈
type Mystack struct {
	queue1 []int
	queue2 []int
}


func Constructor() Mystack {
	return Mystack{
		queue: make([]int,0),
	}
}