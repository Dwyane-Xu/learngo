package main

import "fmt"

// 1049. 最后一块石头的重量 II

func lastStoneWeightII(stones []int) int {
	sum := 0
	length := len(stones)
	for i := 0; i < length; i++ {
		sum += stones[i]
	}

	saveSum := sum
	sum /= 2
	dp := make([]int, sum+1)
	for i := 0; i < length; i++ {
		for j := sum; j >= stones[i]; j-- {
			dp[j] = getMax2(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return saveSum - 2*dp[sum]
}

func getMax2(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
}
