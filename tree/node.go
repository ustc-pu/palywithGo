package main

import (
	"fmt"
)

type treeNode struct {
	value int
	left, right *treeNode
}

func (node treeNode) printNode() {
	fmt.Print(node.value, " ")
}

func createNode(value int) *treeNode {
	return &treeNode{value:value} // can return local variable
}

func (node *treeNode) setValue(val int) {
	if node == nil {
		fmt.Println("nil node")
		return
	}
	node.value = val
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}

	node.left.traverse()
	node.printNode()
	node.right.traverse()
}


func main() {
	var root = treeNode{}
	root = treeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{9, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(1)

	root.printNode()
	fmt.Println()
	if root.right.left.right == nil {
		fmt.Println("nil node")
	}
	root.right.left.setValue(5)
	root.setValue(4)
	root.traverse()
	fmt.Println()
	//
	var pRoot = &treeNode{}
	pRoot.setValue(10)
	pRoot = &root
	//pRoot.traverse()

	nodeSlice := []treeNode {
		{value:-1},
		{},
		{10, nil, &root},
	}
	fmt.Println(nodeSlice)
}
