package main

import "fmt"

// 220. 存在重复元素 III

func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}

	bucket := make(map[int]int, len(nums))
	w := t + 1

	for i := range nums {
		id := keyFunc(nums[i], w)
		if _, ok := bucket[id]; ok {
			return true
		}
		if _, ok := bucket[id-1]; ok && abs(nums[i]-bucket[id-1]) <= t {
			return true
		}
		if _, ok := bucket[id+1]; ok && abs(nums[i]-bucket[id+1]) <= t {
			return true
		}
		bucket[id] = nums[i]
		if i >= k {
			delete(bucket, keyFunc(nums[i-k], w))
		}
	}

	return false
}

func abs(data int) int {
	if data > 0 {
		return data
	} else {
		return -data
	}
}

func keyFunc(num int, w int) int {
	if num >= 0 {
		return num / w
	} else {
		return (num+1)/w - 1
	}
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate2([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate2([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate2([]int{1, 5, 9, 1, 5, 9}, 2, 3))
	fmt.Println(containsNearbyAlmostDuplicate2([]int{-1, -1}, 1, -1))
	fmt.Println(containsNearbyAlmostDuplicate2([]int{2, 1}, 1, 1))
}
