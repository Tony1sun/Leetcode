package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	levl := 0
	queue := make([]*Node, 0)
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			root = queue[0]
			for _, node := range root.Children {
				queue = append(queue, node)
			}
			size--
		}
		levl
	}
}
