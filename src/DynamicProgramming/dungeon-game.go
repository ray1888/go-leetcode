package DynamicProgramming

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}
	row := len(dungeon)
	col := len(dungeon[0])

	dp := make([][]int, row+1)

	for i := 0; i <= row; i++ {
		dp[i] = make([]int, col+1)
		for j := 0; j <= col; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[row][col-1], dp[row-1][col] = 1, 1

	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			minn := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(minn-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}
