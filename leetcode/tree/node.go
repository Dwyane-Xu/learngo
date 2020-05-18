package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}

func (node Node) GetValue() {
	fmt.Printf("%d ", node.Value)
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.GetValue()
	node.Right.Traverse()
}
