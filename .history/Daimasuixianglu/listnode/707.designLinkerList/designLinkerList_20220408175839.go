package main

type MyLinkedList struct {
	head *Node
}

type Node struct {
	Val  int
	Next *Node
	Prev
}