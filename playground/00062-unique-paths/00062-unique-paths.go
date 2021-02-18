package _00062_unique_paths

func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ { // make matrix by m * n
		dp[i] = make([]int, m)
	}

	for i := 0; i < m; i++ { // special case first row
		dp[0][i] = 1
	}
	for i := 0; i < n; i++ { // special case first column
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ { // cross frist
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}
