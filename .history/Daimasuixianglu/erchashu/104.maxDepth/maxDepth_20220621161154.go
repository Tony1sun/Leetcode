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
		
}
