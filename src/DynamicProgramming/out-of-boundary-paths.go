package DynamicProgramming

func findPaths(m int, n int, N int, i int, j int) int {
	if N == 0 {
		return 0
	}
	mod := 1000000007
	var direct = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for k := 0; k < N; k++ {
		tmp := make([][]int, m)
		for i := 0; i < m; i++ {
			tmp[i] = make([]int, n)
		}
		for x := 0; x < m; x++ {
			for y := 0; y < n; y++ {
				for _, d := range direct {
					if x+d[0] < 0 || x+d[0] >= m || y+d[1] < 0 || y+d[1] >= n {
						tmp[x][y]++
					} else {
						tmp[x][y] = (tmp[x][y] + dp[x+d[0]][y+d[1]]) % mod
					}
				}
			}
		}
		dp = tmp
	}
	return dp[i][j]
}
