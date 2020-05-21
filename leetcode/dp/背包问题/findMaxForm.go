package main

import "fmt"

// 474. 一和零
func findMaxForm(strs []string, m int, n int) int {
	length := len(strs)

	nZero := make([]int, length)
	nOne := make([]int, length)
	for i, v1 := range strs {
		for _, v2 := range v1 {
			if v2 == '0' {
				nZero[i]++
			} else {
				nOne[i]++
			}
		}
	}

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < length; i++ {
		for j := m; j >= nZero[i]; j-- {
			for k := n; k >= nOne[i]; k-- {
				dp[j][k] = getMax(dp[j][k], dp[j-nZero[i]][k-nOne[i]]+1)
			}
		}
	}

	return dp[m][n]
}

func getMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
