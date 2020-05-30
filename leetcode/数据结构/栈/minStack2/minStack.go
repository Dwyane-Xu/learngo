package main

import (
	"container/list"
	"fmt"
	"math"
)

type MinStack struct {
	stack    *list.List
	minStack *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	minStack := list.New()
	minStack.PushBack(math.MaxInt32)
	return MinStack{
		stack:    list.New(),
		minStack: minStack,
	}
}

func (this *MinStack) Push(x int) {
	this.stack.PushBack(x)
	v := this.minStack.Back()
	this.minStack.PushBack(min(x, v.Value.(int)))
}

func (this *MinStack) Pop() {
	this.stack.Remove(this.stack.Back())
	this.minStack.Remove(this.minStack.Back())
}

func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}

func (this *MinStack) GetMin() int {
	return this.minStack.Back().Value.(int)
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	obj := Constructor()
	obj.Push(1)
	// obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()
	fmt.Println(param_3)
	fmt.Println(param_4)
}
