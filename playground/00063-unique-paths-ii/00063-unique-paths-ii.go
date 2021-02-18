package _00063_unique_paths_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 { // special case
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ { // make DP matrix
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ { // scan column
		if dp[0][i-1] != 0 && obstacleGrid[0][i] != 1 {
			dp[0][i] = 1
		}
	}
	for i := 1; i < m; i++ { // scan row
		if dp[i-1][0] != 0 && obstacleGrid[i][0] != 1 {
			dp[i][0] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 { // pass obstacleGrid[i][j] == 1
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
