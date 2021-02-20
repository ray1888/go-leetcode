package backpack

import "go-leetcode/src/utils"

func backPackIV(m int, A []int, V []int, times []int) int {
	dp := make([][]int, len(A)+1)
	maxValue := 0
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= len(A); i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			k := m / A[i-1]
			if times[i-1] != 0 {
				k = utils.Min(k, times[i-1])
			}
			for l := 1; l <= k; l++ {
				if j >= (A[i-1] * l) {
					dp[i][j] = utils.Max(dp[i-1][j], dp[i-1][j-(A[i-1]*l)]+V[i-1]*l)
				}
				maxValue = utils.Max(maxValue, dp[i][j])
			}
		}
	}
	return dp[len(A)][m]
}
