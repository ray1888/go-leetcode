package DynamicProgramming

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	length := len(triangle)
	dp := make([]int, length+1)
	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}
	return dp[0]
}
