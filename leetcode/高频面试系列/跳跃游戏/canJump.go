package main

import "fmt"

// 55. 跳跃游戏

func dfs(index int, nums []int, flag []int) int {
	if index == len(nums)-1 {
		return 1
	}
	if nums[index] == 0 {
		return -1
	}
	if flag[index] != 0 {
		return flag[index]
	}

	for i := 1; i <= nums[index]; i++ {
		flag[index] = dfs(index+i, nums, flag)
		if flag[index] == 1 {
			return flag[index]
		}
	}

	return flag[index]
}

func canJump(nums []int) bool {
	if dfs(0, nums, make([]int, len(nums))) == 1 {
		return true
	} else {
		return false
	}
}

func canJump2(nums []int) bool {
	length := len(nums)
	dp := make([]bool, length)
	dp[length-1] = true
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j <= i+nums[i] && j < length; j++ {
			if dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}

func canJump3(nums []int) bool {
	length := len(nums)
	dp := make([]bool, length)
	dp[0] = true
	for i := 0; i < length; i++ {
		if dp[i] {
			for j := i + 1; j <= i+nums[i] && j < length; j++ {
				dp[j] = true
			}
			if dp[length-1] {
				return dp[length-1]
			}
		}
	}
	return dp[length-1]
}

func canJump4(nums []int) bool {
	length := len(nums)
	pos := length - 1
	for i := length - 2; i >= 0; i-- {
		if i+nums[i] >= pos {
			pos = i
		}
	}

	if pos == 0 {
		return true
	} else {
		return false
	}
}

func canJump5(nums []int) bool {
	far := 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > far {
			far = i + nums[i]
		}
		if far == i {
			return false
		}
	}
	return true
}

func main() {
	// fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	// fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump2([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump2([]int{3, 2, 1, 0, 4}))
}
