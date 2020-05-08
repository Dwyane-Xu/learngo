package main

import "fmt"

type NumArray struct {
	value []int
}

func Constructor(nums []int) NumArray {
	numa := NumArray{[]int{0}}
	for i, v := range nums {
		numa.value = append(numa.value, v+numa.value[i])
	}
	return numa
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.value[j+1] - this.value[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}
