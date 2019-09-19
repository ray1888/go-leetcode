package DynamicProgramming

func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxTradeCount := 2
	length := len(prices)
	dp := make([][]int, length*(maxTradeCount+1))
	single_dis := make([]int, length*(maxTradeCount+1))
	for i := 0; i < length; i++ {
		dp[i], single_dis = single_dis[:(maxTradeCount+1)], single_dis[(maxTradeCount+1):]
	}
	minArray := make([]int, length)
	for i := 0; i < length; i++ {
		minArray[i] = prices[0]
	}

	for i := 1; i < length; i++ {
		for j := 1; j <= maxTradeCount; j++ {
			minArray[j] = min(prices[i]-dp[i][j-1], minArray[j])
			dp[i][j] = max(prices[i]-minArray[j], dp[i-1][j])
		}
	}
	return dp[length-1][maxTradeCount]
}
