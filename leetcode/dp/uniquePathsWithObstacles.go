package main

// 63. 不同路径 II

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	row, col := len(obstacleGrid), len(obstacleGrid[0])
	obstacleGrid[0][0] = 1
	for i := 1; i < row; i++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 0 {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}
	for i := 1; i < col; i++ {
		if obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1] == 0 {
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}

	return obstacleGrid[row-1][col-1]
}

func main() {

}
