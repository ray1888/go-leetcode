package DynamicProgramming

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	length := len(grid)
	width := len(grid[0])
	dp := make([][]int, length*width)
	single_dis := make([]int, length*width)
	for i := 0; i < length; i++ {
		dp[i], single_dis = single_dis[:width], single_dis[width:]
	}

	dp[0][0] = grid[0][0]
	l_sum := 0
	for i := 0; i < length; i++ {
		l_sum += grid[i][0]
		dp[i][0] = l_sum
	}

	w_sum := 0
	for j := 0; j < length; j++ {
		w_sum += grid[0][j]
		dp[0][j] = w_sum
	}

	for i := 1; i < length; i++ {
		for j := 1; j < width; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[length-1][width-1]
}
