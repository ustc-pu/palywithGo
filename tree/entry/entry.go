package main

import (
	"fmt"
	"pcheng/playwithGo/tree"
)

type myTreeNode struct{
	node *tree.Node
}

func (aNode *myTreeNode) postOrder() {
	if aNode.node == nil || aNode == nil {
		return
	}

	left := myTreeNode{aNode.node.Left}
	right := myTreeNode{aNode.node.Right}
	left.postOrder()
	right.postOrder()
	aNode.node.PrintNode()
}

func main() {
	var root = tree.Node{}
	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 9}
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
	myNode := myTreeNode{&root}
	myNode.postOrder()
	//
	//var pRoot = &tree.Node{}
	//pRoot.SetValue(10)
	//pRoot = &root
	//pRoot.SetValue(100)
	//pRoot.PrintNode()
	//
	//var nodeSlice = []tree.Node{
	//	{Value: -1},
	//	{},
	//	{10, nil, &root},
	//}
	//fmt.Println(nodeSlice)
	//fmt.Println(root)
}

