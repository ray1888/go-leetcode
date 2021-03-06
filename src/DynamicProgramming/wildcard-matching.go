package DynamicProgramming

func isEqual(sc, pc byte) bool {
	return sc == pc || pc == '?'
}

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for i := 1; i <= m; i++ {
		dp[i][0] = false
	}

	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		} else {
			dp[0][j] = false
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sc := s[i-1]
			pc := p[j-1]
			if isEqual(sc, pc) {
				dp[i][j] = dp[i-1][j-1]
			} else if pc == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[m][n]
}
