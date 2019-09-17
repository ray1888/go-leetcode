package DynamicProgramming

func minmin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

func minDistance(word1 string, word2 string) int {
	if word1 == "" || word2 == "" {
		return 0
	}
	length := len(word1) + 1
	width := len(word2) + 1

	dp := make([][]int, length)
	single_dis := make([]int, (length)*(width))
	for i := 0; i < length; i++ {
		dp[i], single_dis = single_dis[:width], single_dis[width:]
	}

	for i := 0; i < length; i++ {
		dp[i][0] = i
	}
	for j := 0; j < width; j++ {
		dp[0][j] = j
	}

	/*
		here i, j is to fill dp,dp[0][0] mean two input both is "",
		so dp[i][j] map to string compare word1[i-1] and word2[j-1]
	*/
	for i := 1; i < length; i++ {
		for j := 1; j < width; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minmin(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[length-1][width-1]
}
