package main

import (
	"fmt"
	"testing"
)

func Test_MyQueue(t *testing.T) {
	obj := Constructor()
	fmt.Printf("obj= %v\n", obj)
	obj.Push(2)
	fmt.Printf("obj= %v\n", obj)
	obj.Push(10)
	fmt.Printf("obj= %v\n", obj)
	param2 := obj.Pop()
	fmt.Printf("param_2 = %v\n", param2)
	param3 := obj.Peek()
	fmt.Printf("param_3 = %v\n", param3)
	param4 := obj.Empty()
	fmt.Printf("param_4 = %v\n", param4)
	obj.Push(3)
	fmt.Printf()
}
