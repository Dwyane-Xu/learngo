package main

import "fmt"

// 219. 存在重复元素 II

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]struct{})
	for i, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
		if len(seen) > k {
			delete(seen, nums[i-k])
		}
	}
	return false
}

func main() {
	// fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	// fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
