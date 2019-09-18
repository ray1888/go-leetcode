package DynamicProgramming

import "math"

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}
	length := len(coins) + 1
	dp := make([][]int, length*(amount+1))
	single_dis := make([]int, length*(amount+1))
	for i := 0; i < length; i++ {
		dp[i], single_dis = single_dis[:(amount+1)], single_dis[(amount+1):]
	}

	for i := 0; i < length; i++ {
		dp[i][0] = 0
	}

	for j := 0; j <= amount; j++ {
		dp[0][j] = int(math.Pow(2, 32))
	}

	for i := 1; i < length; i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			} else {
				dp[i][j] = min(dp[i-1][j], int(math.Pow(2, 32)))
			}
		}
	}
	// this judgement is for leetcode testcase passed
	if dp[length-1][amount] == int(math.Pow(2, 32)) {
		return -1
	} else {
		return dp[length-1][amount]
	}

}
