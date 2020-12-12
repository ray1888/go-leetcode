package DynamicProgramming

func backPackIII(m int, A []int, V []int) int {
	dp := make([][]int, len(A)+1)
	maxValue := 0
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= len(A); i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			k := m / A[i-1]
			for l := 1; l <= k; l++ {
				if j >= A[i-1]*k {
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-A[i-1]]+V[i-1]*k)
				}
				maxValue = max(maxValue, dp[i][j])
			}
		}
	}
	return dp[len(A)][m]
}

func backPackIIIOnSpace(m int, A []int, V []int) int {
	dp := make([]int, m+1)

	for i := 1; i <= len(A); i++ {
		for j := A[i-1]; j <= m; j++ {
			dp[j] = max(dp[j], dp[j-A[i-1]]+V[i-1])
		}
	}
	return dp[m]
}
