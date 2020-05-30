package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	maxNum, nowNum, digit := 0, 0, -1
	var result []int
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Val == digit {
			nowNum++
		} else {
			digit = root.Val
			nowNum = 0
		}

		if nowNum == maxNum {
			result = append(result, root.Val)
		} else if nowNum > maxNum {
			result = []int{root.Val}
			maxNum = nowNum
		}

		root = root.Right
	}
	return result
}

func main() {

}
