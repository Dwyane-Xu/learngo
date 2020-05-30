package main

import (
	"container/list"
	"fmt"
	"math"
)

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

type numbers struct {
	visit  bool
	indexs *list.List
}

func main() {
	// numMap := make(map[int]numbers, 5)
	// for i, v := range []int{1, 2, 4, 4, 5} {
	// 	if _, ok := numMap[v]; !ok {
	// 		numMap[v] = numbers{false, list.New()}
	// 	}
	// 	numMap[v].indexs.PushBack(i)
	// }
	// fmt.Println(numMap)

	// numMap := make(map[int]numbers, 10)
	// numMap[0] = numbers{false, list.New()}
	// numMap[0].indexs.PushBack(1)
	// fmt.Println(numMap[0].indexs.Back().Value)
	// numMap[0] = numbers{}

	// numMap := make(map[int]*list.List, 10)
	// numMap[0] = list.New()
	// numMap[0].PushBack(1)
	// fmt.Println(numMap[0].Back().Value)

	seen := make(map[int]bool, 2)
	fmt.Println(seen[2])
}
