package DynamicProgramming

func changerecur(coins []int, start, sum int) int {
	if sum == 0 {
		return 1
	}
	if sum < 0 {
		return 0
	}
	result := 0
	for i := start; i < len(coins); i++ {
		result += changerecur(coins, i, sum-coins[i])
	}
	return result
}

func ChangeRecursive(amount int, coins []int) int {
	if amount == 0 || len(coins) == 0 {
		return 0
	}
	return changerecur(coins, 0, amount)
}

func ChangeDP(amount int, coins []int) int {
	if amount == 0 || len(coins) == 0 {
		return 0
	}
	length := len(coins)
	dp := make([][]int, (length + 1))
	single_dis := make([]int, (length+1)*(amount+1))
	for i := 0; i <= length; i++ {
		dp[i], single_dis = single_dis[:(amount+1)], single_dis[(amount+1):]
	}

	for i := 0; i <= length; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= length; i++ {
		for j := 1; j <= amount; j++ {
			userCurCoin := 0
			if j >= coins[i-1] {
				userCurCoin = dp[i][j-coins[i-1]]
			}
			dp[i][j] = dp[i-1][j] + userCurCoin
		}
	}
	return dp[length][amount]
}
