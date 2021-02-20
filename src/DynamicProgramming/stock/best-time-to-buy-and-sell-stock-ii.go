package stock

func maxProfitII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	result := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			result += (prices[i+1] - prices[i])
		}
	}
	return result
}
