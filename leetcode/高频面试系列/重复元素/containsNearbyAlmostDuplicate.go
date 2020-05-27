package main

import (
	"fmt"
	"math"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gutil"
)

// 220. 存在重复元素 III

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) == 0 {
		return false
	}

	minMap := gmap.NewTreeMap(gutil.ComparatorInt)
	maxMap := gmap.NewTreeMap(gutil.ComparatorInt)
	minMap.Set(nums[0], 1)
	maxMap.Set(nums[0], 1)

	for i := 1; i < len(nums); i++ {
		if i > k {
			remove(minMap, nums[i-k-1])
			remove(maxMap, nums[i-k-1])
		}
		if minMap.Size() == 0 {
			continue
		}

		minNum := minMap.Left().Key.(int)
		maxNum := maxMap.Right().Key.(int)
		if nums[i] >= maxNum {
			if nums[i]-maxNum <= t {
				return true
			}
		} else if nums[i] <= minNum {
			if minNum-nums[i] <= t {
				return true
			}
		} else {
			for j := 1; j <= k; j++ {
				if i-j < 0 {
					break
				}
				if math.Abs(float64(nums[i]-nums[i-j])) <= float64(t) {
					return true
				}
			}
		}

		add(minMap, nums[i])
		add(maxMap, nums[i])
	}

	return false
}

func add(treeMap *gmap.TreeMap, num int) {
	v := treeMap.Get(num)
	if v == nil {
		treeMap.Set(num, 1)
	} else {
		treeMap.Set(num, v.(int)+1)
	}
}

func remove(treeMap *gmap.TreeMap, num int) {
	v := treeMap.Get(num)
	if v == 1 {
		treeMap.Remove(num)
	} else {
		treeMap.Set(num, v.(int)-1)
	}
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}
