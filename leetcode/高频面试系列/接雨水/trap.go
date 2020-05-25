package main

import "fmt"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	length := len(height)
	l, r := 1, length-2
	lMax, rMax, result := height[0], height[length-1], 0
	for l <= r {
		lMax = getMax(lMax, height[l])
		rMax = getMax(rMax, height[r])
		if lMax < rMax {
			result += lMax - height[l]
			l++
		} else {
			result += rMax - height[r]
			r--
		}
	}

	return result
}

func getMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
