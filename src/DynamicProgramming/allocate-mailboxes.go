package DynamicProgramming

import (
	"math"
	"sort"
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func initMatrix(m, n int, value int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	if value != 0 {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = value
			}
		}
	}
	return matrix
}

/*
	此题可以使用DP作为代表i位置邮箱放置的地方，k代表第几个邮箱
	递推状态方程为：
	就是对比前i-1 个邮箱的距离+ 放在当前i到达其他的中位数距离最小
*/

func minDistanceMailBox(houses []int, k int) int {
	if len(houses) == 0 || k == 0 {
		return 0
	}
	sort.Ints(houses)
	n := len(houses)
	// 求得中位数
	medSum := initMatrix(n, n, 0)
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			medSum[i][j] = medSum[i+1][j-1] + houses[j] - houses[i]
		}
	}

	dp := initMatrix(n, k+1, math.MaxInt32)
	for i := 0; i < n; i++ {
		dp[i][1] = medSum[0][i]
		for j := 2; j <= k && j <= i+1; j++ {
			for i0 := 0; i0 < i; i0++ {
				dp[i][j] = min(dp[i][j], dp[i0][j-1]+medSum[i0+1][i])
			}
		}
	}
	return dp[n-1][k]
}
