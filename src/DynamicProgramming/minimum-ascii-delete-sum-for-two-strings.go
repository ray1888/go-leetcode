package DynamicProgramming

func minimumDeleteSum(s1 string, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	l1 := len(s1)
	l2 := len(s2)

	dp := make([][]int, l1+1)

	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := l1 - 1; i >= 0; i-- {
		dp[i][l2] = dp[i+1][l2] + int(s1[i])
	}

	for i := l2 - 1; i >= 0; i-- {
		dp[l1][i] = dp[l1][i+1] + int(s2[i])
	}

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = min(dp[i+1][j]+int(s1[i]), dp[i][j+1]+int(s2[j]))
			}
		}
	}
	return dp[0][0]
}
