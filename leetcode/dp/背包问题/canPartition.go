package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < length; i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[sum]
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
	fmt.Println(canPartition([]int{1}))
	fmt.Println(canPartition([]int{23, 13, 11, 7, 6, 5, 5}))
}
