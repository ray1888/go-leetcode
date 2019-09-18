package DynamicProgramming

func fibRecur(N int) int {
	if N <= 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fibRecur(N-1) + fibRecur(N-2)
}

func fibDP(N int) int {
	if N == 0 {
		return 0
	} else if N <= 2 {
		return 1
	}
	dp := make([]int, N)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N-1]
}
