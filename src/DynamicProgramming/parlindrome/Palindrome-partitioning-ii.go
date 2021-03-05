package parlindrome

import (
	"go-leetcode/src/utils"
	"math"
)

func minCut(s string) int {

	if s == "" {
		return 0
	}
	length := len(s)
	dp := make([][]bool, length*length)
	single_dis_dp := make([]bool, length*length)
	for i := 0; i < length; i++ {
		dp[i], single_dis_dp = single_dis_dp[:length], single_dis_dp[length:]
	}

	minArray := make([]int, length)
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if s[i] == s[j] && (i-j < 2 || dp[i-1][j+1]) {
				dp[i][j] = true
			}
		}
	}
	IntMax := int(math.Pow(2, 32))

	for i := 1; i < length; i++ {
		tmp := IntMax
		for j := 0; j <= i; j++ {
			if dp[j][i] {
				if j == 0 {
					tmp = 0
					break
				} else {
					tmp = utils.Min(tmp, minArray[i-1]+1)
				}
			}
		}
		minArray[i] = tmp
	}
	return minArray[length-1]
}

func dp(s string) int {
	if len(s) == 0 {
		return -1
	}
	n := len(s)

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = true
			} else if i+1 == j {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
		}
	}

	cut := make([]int, n+1)
	cut[0] = -1
	for j := 0; j < n; j++ {
		cut[j+1] = cut[j] + 1
		for i := j - 1; i >= 0; i-- {
			if dp[i][j] {
				cut[j+1] = utils.Min(cut[j+1], cut[i]+1)
			}
		}
	}

	return cut[n]
}
