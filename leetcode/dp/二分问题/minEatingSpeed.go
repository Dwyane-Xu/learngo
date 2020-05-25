package main

import "fmt"

// 875. 爱吃香蕉的珂珂

func minEatingSpeed(piles []int, H int) int {
	l, r := 1, 0
	for i := 0; i < len(piles); i++ {
		r = getMax(r, piles[i])
	}

	for l < r {
		mid := (l + r) >> 1
		if eat(piles, H, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
	// return l
}

func eat(piles []int, H int, k int) bool {
	now := 0
	for i := 0; i < len(piles); i++ {
		now += (piles[i] + k - 1) / k
		if now > H {
			return false
		}
	}
	return true
}

func getMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}
