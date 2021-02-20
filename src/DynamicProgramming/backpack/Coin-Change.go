package backpack

import (
	"go-leetcode/src/utils"
	"math"
)

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}
	length := len(coins) + 1
	dp := make([][]int, length*(amount+1))
	singleDis := make([]int, length*(amount+1))
	for i := 0; i < length; i++ {
		dp[i], singleDis = singleDis[:(amount+1)], singleDis[(amount+1):]
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
				dp[i][j] = utils.Min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			} else {
				dp[i][j] = utils.Min(dp[i-1][j], int(math.Pow(2, 32)))
			}
		}
	}
	// this judgement is for leetcode testcase passed
	if dp[length-1][amount] == int(math.Pow(2, 32)) {
		return -1
	}
	return dp[length-1][amount]
}

func coinChangeOn(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			var usedCoin int
			if j < coins[i-1] {
				usedCoin = math.MaxInt32
			} else {
				usedCoin = dp[j-coins[i-1]]
			}
			if usedCoin != math.MaxInt32 {
				usedCoin++
			}
			dp[j] = utils.Min(usedCoin, dp[j])
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
