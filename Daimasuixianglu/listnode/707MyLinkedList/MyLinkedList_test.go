package main

import (
	"fmt"
	"testing"
)

func Test_MyLinkedList(t *testing.T) {
	obj := Constructor()
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtHead(58)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtIndex(1, 27)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	param1 := obj.Get(1)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	obj.AddAtHead(38)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtHead(45)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtTail(36)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.AddAtIndex(1, 24)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.DeleteAtIndex(3)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))

}

func MList2Ints(this *MyLinkedList) []int {
	res := []int{}
	cur := this.head
	for cur != nil {
		res = append(res, cur.val)
		cur = cur.next
	}
	return res
}
