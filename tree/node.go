package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left, Right *Node
}

func (node Node) PrintNode() {
	fmt.Print(node.Value, " ")
}

func CreateNode(value int) *Node {
	return &Node{Value:value} // can return local variable
}

func (node *Node) SetValue(val int) {
	if node == nil {
		fmt.Println("nil node")
		return
	}
	node.Value = val
}





