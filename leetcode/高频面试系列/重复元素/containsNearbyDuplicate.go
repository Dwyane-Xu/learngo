package main

import "fmt"

// 219. 存在重复元素 II

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]bool)
	for i, v := range nums {
		if mp[v] {
			return true
		}
		mp[v] = true
		if len(mp) > k {
			delete(mp, nums[i-k])
		}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
