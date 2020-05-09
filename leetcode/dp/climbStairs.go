package main

import "fmt"

// 70. 爬楼梯

// recursion
func climbStairsWithRecursion(i, n int, flag map[int]int) int {
	if i > n {
		return 0
	} else if i == n {
		return 1
	} else if flag[i] > 0 {
		return flag[i]
	}

	flag[i] = climbStairsWithRecursion(i+1, n, flag) + climbStairsWithRecursion(i+2, n, flag)
	return flag[i]
}

// dp
func climbStairsWithDp(n int) int {
	one, two := 1, 1
	for i := 1; i < n; i++ {
		one, two = two, one+two
	}
	return two
}

func climbStairs(n int) int {
	// return climbStairsWithRecursion(0, n, map[int]int{})
	return climbStairsWithDp(n)
}

func main() {
	fmt.Println(climbStairs(2))
}
