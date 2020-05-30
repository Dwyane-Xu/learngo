package main

import "fmt"

// 1306. 跳跃游戏 III

func canReach(arr []int, start int) bool {
	return dfs2(arr, start, make([]bool, len(arr)))
}

func dfs2(arr []int, index int, flag []bool) bool {
	if arr[index] == 0 {
		return true
	}
	if flag[index] {
		return false
	}

	flag[index] = true
	var left, right bool
	if index-arr[index] >= 0 {
		left = dfs2(arr, index-arr[index], flag)
	}
	if index+arr[index] < len(arr) {
		right = dfs2(arr, index+arr[index], flag)
	}

	return left || right
}

func canReach2(arr []int, start int) bool {
	length := len(arr)
	stack := make([]int, length)
	ma := make(map[int]bool, length)
	stack[0] = start
	ma[start] = true
	top := 0
	for top != -1 {
		index := stack[top]
		top--
		if arr[index] == 0 {
			return true
		}
		left, right := index-arr[index], index+arr[index]
		if left >= 0 && !ma[left] {
			top++
			stack[top] = left
			ma[left] = true
		}
		if right < length {
			top++
			stack[top] = right
			ma[right] = true
		}
	}
	return false
}

func main() {
	// fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
	// fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0))
	// fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2))

	fmt.Println(canReach2([]int{4, 2, 3, 0, 3, 1, 2}, 5))
	fmt.Println(canReach2([]int{4, 2, 3, 0, 3, 1, 2}, 0))
	fmt.Println(canReach2([]int{3, 0, 2, 1, 2}, 2))
}
