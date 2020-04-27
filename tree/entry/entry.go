package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	*tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.GetValue()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This method is shadowed.")
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = tree.CreateTreeNode(4)
	root.Right = tree.CreateTreeNode(5)
	root.Left.Right = tree.CreateTreeNode(1)
	root.Right.Left = tree.CreateTreeNode(2)

	root.Left.Right.GetValue()
	fmt.Println()
	root.Left.Right.SetValue(6)
	root.Left.Right.GetValue()
	fmt.Println()

	root.Traverse()
	root.Node.Traverse()
	fmt.Println()
	root.postOrder()
	fmt.Println()
}
