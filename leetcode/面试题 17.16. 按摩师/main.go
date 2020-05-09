package main

import (
	"fmt"
)

func massage(nums []int) int {
	l1, l2, max := 0, 0, 0
	for _, v := range nums {
		if l1+v > l2 {
			max = l1 + v
		} else {
			max = l2
		}
		l1 = l2
		l2 = max
	}
	return max
}

func main() {
	fmt.Println(massage([]int{1, 2, 3, 1}))
	fmt.Println(massage([]int{2, 7, 9, 3, 1}))
	fmt.Println(massage([]int{2, 1, 4, 5, 3, 1, 1, 3}))
}
