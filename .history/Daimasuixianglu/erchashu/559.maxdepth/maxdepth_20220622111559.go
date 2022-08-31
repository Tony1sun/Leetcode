package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {