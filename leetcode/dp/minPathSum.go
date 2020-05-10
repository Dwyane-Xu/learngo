package main

// 64. 最小路径和

// use original two-dimensional array
func minPathSumOne(grid [][]int) int {
	for i := range grid {
		for j := range grid[i] {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = getMin(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

// user new one-dimensional array
func minPathSumTwo(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
			} else if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = getMin(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[len(grid[0])-1]
}

func main() {

}
