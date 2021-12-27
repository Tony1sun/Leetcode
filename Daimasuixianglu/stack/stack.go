package main

import (
	"errors"
	"fmt"
)

// 使用数组模拟一个栈
type Stack struct {
	MaxTop int    // 最大存放个数
	Top    int    // 栈顶
	arr    [5]int //数组模拟栈
}

func (this *Stack) Push(val int) (err error) {
	// 判断栈是否已满
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	// 放入数据
	this.arr[this.Top] = val
	return
}

// 遍历栈，从栈顶开始遍历
func (this *Stack) List() {
	// 判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下:")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	// 判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	// 先取值，再this.Top--
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	//入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	//显示
	stack.List()

	val, _ := stack.Pop()
	fmt.Println("出栈val=", val)
	stack.List()

	fmt.Println()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	fmt.Println("出栈val=", val)

	stack.List()

}
