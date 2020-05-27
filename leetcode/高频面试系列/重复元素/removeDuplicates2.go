package main

import "fmt"

// 80. 删除排序数组中的重复项 II

func removeDuplicates2(nums []int) int {
	// slow, fast := 0, 1
	// count := 1
	// for fast < len(nums) {
	// 	if nums[slow] != nums[fast] {
	// 		slow += count
	// 		nums[slow] = nums[fast]
	// 		count = 1
	// 	} else if count == 1 {
	// 		nums[slow+1] = nums[fast]
	// 		count++
	// 	}
	// 	nums[slow+count] = nums[fast]
	// 	fast++
	// }
	// return slow + count

	j, count := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func main() {
	fmt.Println(removeDuplicates2([]int{1, 1, 1, 2, 2, 3}))
	a := []int{1, 1, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates2(a))
	fmt.Println(a)
}
