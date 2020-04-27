package main

import (
	"fmt"
	"math"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || numRows == len(s) {
		return s
	}

	var result strings.Builder
	n := len(s)

	add := 2*numRows - 2
	for i := 0; i < n; i += add {
		result.WriteByte(s[i])
	}

	add = 2*numRows - 4
	for row := 1; row < numRows/2; row++ {
		for i := row; i < n; i = i + add + 2*row {
			result.WriteByte(s[i])
			if i+add < n {
				result.WriteByte(s[i+add])
			}
		}
		add -= 2
	}

	if numRows%2 == 1 {
		add = numRows - 1
		for i := numRows / 2; i < n; i += add {
			result.WriteByte(s[i])
		}
	}

	if numRows%2 == 0 {
		add = numRows - 2
	} else {
		add = numRows - 3
	}
	for row := (numRows + 1) / 2; row < numRows-1; row++ {
		for i := row; i < n; i = i + add + 2*row {
			result.WriteByte(s[i])
			if i+add < n {
				result.WriteByte(s[i+add])
			}
		}
		add -= 2
	}

	add = 2*numRows - 2
	for i := numRows - 1; i < n; i += add {
		result.WriteByte(s[i])
	}

	return result.String()
}

func convertTwo(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var res string
	var tmp = make([]string, numRows)
	curPos := 0
	shouldTurn := -1
	for _, val := range s {
		tmp[curPos] += string(val)
		if curPos == 0 || curPos == numRows-1 {
			shouldTurn = -shouldTurn
		}
		curPos += shouldTurn
	}

	for _, val := range tmp {
		res += val
	}
	return res
}

func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x /= 10
		if result < math.MinInt32 || result > math.MaxInt32 {
			return 0
		}
	}
	return result
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	z := x
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x /= 10
	}

	return z == y
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxHeight := 0
	for l < r {
		maxHeight = getMax(maxHeight, (r-l+1)*getMin(height[l], height[r]))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return maxHeight
}

func getMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func getMin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func findRepeatedDnaSequences(s string) []string {
	res := make(map[string]bool)
	set := make(map[string]bool)
	for i := 0; i <= len(s)-10; i++ {
		key := s[i : i+10]
		if set[key] {
			res[key] = true
		} else {
			set[key] = true
		}
	}

	ret := make([]string, 0)
	for k := range res {
		ret = append(ret, k)
	}

	return ret
}

func hammingWeight(num uint32) int {
	var bits, mask uint32 = 0, 1
	for i := 0; i < 32; i++ {
		if (num & mask) != 0 {
			bits++
		}
		mask <<= 1
	}
	return int(bits)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	levelOrder(root, 0, &res)
	return res
}

func levelOrder(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}
	if level == len(*res) {
		*res = append(*res, root.Val)
	}
	levelOrder(root.Right, level+1, res)
	levelOrder(root.Left, level+1, res)
}

func main() {
	left := TreeNode{Val: 2}
	right := TreeNode{Val: 3}
	root := TreeNode{Val: 1, Left: &left, Right: &right}
	fmt.Println(rightSideView(&root))
}
