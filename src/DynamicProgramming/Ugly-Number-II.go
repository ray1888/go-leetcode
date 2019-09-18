package DynamicProgramming

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	} else if n < 1 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = 1
	p2 := 0
	p3 := 0
	p5 := 0
	for i := 1; i < n; i++ {
		uglyNum := min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
		dp[i] = uglyNum
		if uglyNum%2 == 0 {
			p2 += 1
		}
		if uglyNum%3 == 0 {
			p3 += 1
		}
		if uglyNum%5 == 0 {
			p5 += 1
		}
	}
	return dp[n-1]
}
