package main

import "fmt"

// 217. 存在重复元素

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, v := range nums {
		if !set[v] {
			set[v] = true
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
}
