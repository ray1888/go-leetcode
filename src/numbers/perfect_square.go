package numbers

import "math"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func numSquaresDP(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func isSquareBool(n int) bool {
	a := int(math.Sqrt(float64(n)))
	return a*a == n
}

func numSquaresMath(n int) int {
	if isSquareBool(n) {
		return 1
	}
	for i := 1; i*i <= n; i++ {
		if isSquareBool(n - i*i) {
			return 2
		}
	}
	for n%4 == 0 {
		n /= 4
	}
	if n%8 != 7 {
		return 3
	}
	return 4
}
