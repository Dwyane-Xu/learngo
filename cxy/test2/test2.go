package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	x, y int
}

var p = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func numIslands(grid [][]byte) int {
	res := 0
	for i := range grid {
		for j, v := range grid[i] {
			if v == '1' {
				bfs(grid, i, j)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	for i := 0; i < 4; i++ {
		dfs(grid, x+p[i][0], y+p[i][1])
	}
}

func bfs(grid [][]byte, x, y int) {
	queue := list.New()
	queue.PushBack(Node{x, y})
	grid[x][y] = '1'
	for queue.Len() != 0 {
		node := queue.Remove(queue.Front()).(Node)
		for i := 0; i < 4; i++ {
			nx, ny := node.x+p[i][0], node.y+p[i][1]
			if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) || grid[nx][ny] == '0' {
				continue
			}
			grid[nx][ny] = '0'
			queue.PushBack(Node{nx, ny})
		}
	}
}

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '1'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}))
}
