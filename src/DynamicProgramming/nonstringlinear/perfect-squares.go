package nonstringlinear

import (
	"go-leetcode/src/utils"
	"math"
)

func numSquares(n int) int {

	squares := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for _, square := range squares {
			if i < square {
				break
			}
			dp[i] = utils.Min(dp[i], dp[i-square]+1)
		}
	}
	return dp[n]
}
