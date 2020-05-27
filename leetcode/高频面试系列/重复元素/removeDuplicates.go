package main

import (
	"fmt"
)

// 26. 删除排序数组中的重复项

func removeDuplicates(nums []int) int {
	// length := len(nums)
	// slow, fast := 0, 1
	// for fast < length {
	// 	if nums[slow] != nums[fast] {
	// 		slow++
	// 		nums[slow] = nums[fast]
	// 	}
	// 	fast++
	// }
	// return slow + 1

	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
