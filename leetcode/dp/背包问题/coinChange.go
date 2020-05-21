package main

import "fmt"

func coinChange(coins []int, amount int) int {
	length := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 0; i < length; i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = getMin(dp[j], dp[j-coins[i]]+1)
		}
	}

	if dp[amount] != amount+1 {
		return dp[amount]
	} else {
		return -1
	}
}

func getMin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
}
