package main

type treenode struct {
	val   int
	left  *treenode
	right *treenode
}

func maxdepth(root *treenode) int {
	levl := 0
	queue := make([]*treenode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			node := queue[0]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				}
		}
	}
}
