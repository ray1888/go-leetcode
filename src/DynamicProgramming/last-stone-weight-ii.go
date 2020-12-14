package DynamicProgramming

func lastStoneWeightII(stones []int) int {
	// m := len(stones)
	sum := 0
	bagsize := 0
	for _, item := range stones {
		sum += item
	}
	bagsize = sum / 2
	dp := make([]int, bagsize+1)

	for i := 1; i <= len(stones); i++ {
		for j := bagsize; j >= stones[i-1]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i-1]]+stones[i-1])
		}
	}
	return sum - 2*dp[bagsize]
}
