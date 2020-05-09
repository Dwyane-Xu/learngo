package main

import "fmt"

// 746. 使用最小花费爬楼梯

func minCostClimbingStairs(cost []int) int {
	one, two := 0, 0
	for i := 0; i < len(cost); i++ {
		one, two = two, getMin(one, two)+cost[i]
	}
	return getMin(one, two)
}

func getMin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
