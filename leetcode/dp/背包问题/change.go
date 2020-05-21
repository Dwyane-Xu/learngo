package main

import "fmt"

// 322. 零钱兑换 II

func change(amount int, coins []int) int {
	length := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < length; i++ {
		for j := 1; j <= amount; j++ {
			if coins[i] <= j {
				dp[j] += dp[j-coins[i]]
			}
		}
	}

	return dp[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
