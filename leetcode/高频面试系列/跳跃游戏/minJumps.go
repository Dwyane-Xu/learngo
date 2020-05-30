package main

import (
	"container/list"
	"fmt"
)

// 1345. 跳跃游戏 IV

type point struct {
	index int
	count int
}

func minJumps(arr []int) int {
	length := len(arr)
	numMap := make(map[int][]int, length)
	for i, v := range arr {
		numMap[v] = append(numMap[v], i)
	}

	visit := make(map[int]bool, length)
	visit[0] = true
	queue := list.New()
	queue.PushBack(point{0, 0})
	for queue.Len() != 0 {
		curPoint := queue.Front()
		queue.Remove(curPoint)
		index := curPoint.Value.(point).index
		count := curPoint.Value.(point).count
		if index == length-1 {
			return count
		}

		if index+1 < length && !visit[index+1] {
			visit[index+1] = true
			queue.PushBack(point{index + 1, count + 1})
		}
		if index-1 >= 0 && !visit[index-1] {
			visit[index-1] = true
			queue.PushBack(point{index - 1, count + 1})
		}

		value := arr[index]
		for _, in := range numMap[value] {
			if (in < index-1 || in > index+1) && !visit[in] {
				visit[in] = true
				queue.PushBack(point{in, count + 1})
			}
		}
		delete(numMap, value)
	}

	return -1
}

func main() {
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
	fmt.Println(minJumps([]int{7}))
	fmt.Println(minJumps([]int{7, 6, 9, 6, 9, 6, 9, 7}))
	fmt.Println(minJumps([]int{6, 1, 9}))
	fmt.Println(minJumps([]int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13}))
}
