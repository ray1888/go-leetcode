package DynamicProgramming

//双串问题
//此处为双串问题，因为维度直接是两个字符串
//dp[i][j] 表达的是，当第一个字符串到i的长度；和第二个字符串到j的长度的时候，公共的字符串有多长

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	l1 := len(text1)
	l2 := len(text2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}
