package DynamicProgramming

// 本体的核心逻辑在于，可以通过遍历数组映射为一个DP的矩阵，
// 然后通过矩阵的变换，（从左上角到右下角，计算出来最长的边长是多少）
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	row := len(matrix)
	col := len(matrix[0])
	maxSide := 0
	dp := make([][]int, row)

	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		for j := 0; j < col; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}
