package main

import "fmt"

// 300. 最长上升子序列

func lengthOfLISOne(nums []int) int {
	dp := make([]int, len(nums))
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = getMax(dp[i], dp[j]+1)
			}
		}
		maxLen = getMax(maxLen, dp[i])
	}

	return maxLen
}

func getMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func lengthOfLISTwo(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxLen := 1
	flag := make([]int, len(nums))
	flag[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > flag[maxLen-1] {
			flag[maxLen] = nums[i]
			maxLen++
		} else {
			l, r := 0, maxLen-1
			for l <= r {
				mid := (l + r) >> 1
				if nums[i] > flag[mid] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			flag[l] = nums[i]
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLISOne([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLISTwo([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
