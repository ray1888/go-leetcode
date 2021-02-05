package stock

import (
	"go-leetcode/src/utils"
	"math"
)

func maxProfitIV(K int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)

	dp := make([][][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([][]int, K+1)

		for j := 0; j <= K; j++ {
			dp[i][j] = make([]int, 2)
		}
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt32
	}

	for i := 0; i < n; i++ {
		for k := K; k >= 1; k-- {
			if (i - 1) == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = utils.Max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = utils.Max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][K][0]
}
