package nonstringlinear

import "go-leetcode/src/utils"

/*
	此题思路：
	本质上把每个数i最大能组成的乘积放入到DP数组中，然后通过遍历到n，来做差，是使用DP的最大数组乘积来乘法处理
*/

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = utils.Max(curMax, utils.Max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}
