package parlindrome

import "go-leetcode/src/utils"

func longestPalindromeSubseq(s string) int {
	if s == "" || len(s) == 0 {
		return 0
	}
	length := len(s)
	dp := make([][]int, length)
	single_dis := make([]int, length*length)

	for i := 0; i < length; i++ {
		// here dp mean s from index i to index j is Palindromic string with bool result
		dp[i], single_dis = single_dis[:length], single_dis[length:]
	}

	for i := length - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = utils.Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][length-1]
}
