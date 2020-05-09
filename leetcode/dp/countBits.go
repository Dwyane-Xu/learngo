package main

import "fmt"

// 338. 比特位计数

// my way
func countBitsOne(num int) []int {
	if num == 0 {
		return []int{0}
	}

	n, max := num, 0
	for n > 0 {
		n /= 2
		max++
	}

	dp := make([]int, 1<<uint(max))
	dp[0] = 0

	begin := 1
	for i := 1; i <= max; i++ {
		for j := begin; j < begin+begin/2; j++ {
			dp[j] = dp[j-begin/2]
		}
		for j := begin + begin/2; j < begin*2; j++ {
			dp[j] = dp[j-begin/2] + 1
		}
		begin *= 2
	}

	return dp[0 : num+1]
}

// the most significant bit
func countBitsTwo(num int) []int {
	dp := make([]int, num+1)
	i, m := 0, 1
	for m <= num {
		for i < m && i+m <= num {
			dp[i+m] = dp[i] + 1
			i++
		}
		i = 0
		m <<= 1
	}
	return dp
}

// the less significant bit
func countBitsThree(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}
	return dp
}

// last set bit
func countBitsFour(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}
	return dp
}

func main() {
	fmt.Println(countBitsFour(8))
}
