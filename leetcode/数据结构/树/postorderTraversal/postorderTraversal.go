package main

// 145. 二叉树的后序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalOne(root *TreeNode) []int {
	var res []int
	postOrder(root, &res)
	return res
}

func postOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	postOrder(root.Left, res)
	postOrder(root.Right, res)
	*res = append(*res, root.Val)
}

func postorderTraversalTwo(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	for i := 0; i < len(res)/2; i++ {
		j := len(res) - i - 1
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func postorderTraversalThree(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for len(stack) != 0 || root != nil {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Right
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Left
		}
	}

	for i := 0; i < len(res)/2; i++ {
		j := len(res) - i - 1
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func main() {

}
