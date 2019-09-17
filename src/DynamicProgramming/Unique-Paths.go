package DynamicProgramming

func uniquePathsDP(m int, n int) int {
	dp := make([][]int, m*n)
	single_dis := make([]int, m*n)
	for i := 0; i < m; i++ {
		dp[i], single_dis = single_dis[:n], single_dis[n:]
	}

	dp[0][0] = 1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
