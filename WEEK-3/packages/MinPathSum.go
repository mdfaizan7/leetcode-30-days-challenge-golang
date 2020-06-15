package packages

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	dp := make([][]int, len(grid))
	for i := range grid {
		dp[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] += grid[i][j]

			if i > 0 && j > 0 {
				dp[i][j] += min(dp[i-1][j], dp[i][j-1])
			} else if i > 0 {
				dp[i][j] += dp[i-1][j]
			} else if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

// TestMinPathSum is testing function
func TestMinPathSum() {
	t1 := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	t2 := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("\nTesting for Min Path Sum: %t %t \n", minPathSum(t1) == 7, minPathSum(t2) == 12)
}
