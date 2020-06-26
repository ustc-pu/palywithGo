package main

import (
	"fmt"
	"pcheng/learngo/.idea/tree"
)

func main() {
	var root = tree.Node{}
	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{9, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(1)

	root.PrintNode()
	fmt.Println()
	if root.Right.Left.Right == nil {
		fmt.Println("nil node")
	}
	root.Right.Left.SetValue(5)
	root.SetValue(4)
	root.Traverse()
	fmt.Println()
	//
	var pRoot = &tree.Node{}
	pRoot.SetValue(10)
	pRoot = &root
	pRoot.SetValue(100)
	pRoot.PrintNode()

	nodeSlice := []tree.Node {
		{Value:-1},
		{},
		{10, nil, &root},
	}
	fmt.Println(nodeSlice)
	fmt.Println(root)
}

