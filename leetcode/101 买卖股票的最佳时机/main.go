package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	min, max := math.MaxInt32, 0
	for _, v := range prices {
		if max < v-min {
			max = v - min
		}
		if min > v {
			min = v
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
