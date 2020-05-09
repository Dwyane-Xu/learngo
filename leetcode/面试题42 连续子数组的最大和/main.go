package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	sum, max := 0, math.MinInt32
	for _, v := range nums {
		if sum > 0 {
			sum = sum + v
		} else {
			sum = v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
