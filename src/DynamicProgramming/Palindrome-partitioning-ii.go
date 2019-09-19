package DynamicProgramming

import "math"

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
					tmp = min(tmp, minArray[i-1]+1)
				}
			}
		}
		minArray[i] = tmp
	}
	return minArray[length-1]
}
