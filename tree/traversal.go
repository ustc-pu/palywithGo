package tree

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.PrintNode()
	node.Right.Traverse()
}
