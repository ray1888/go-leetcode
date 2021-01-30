package DynamicProgramming

import (
	"fmt"
	"math"
)

func largestSumOfAveragesSpaceOnSquare(A []int, K int) float64 {
	var sum float64
	if K == 1 {
		for i := 1; i <= len(A); i++ {
			sum += float64(A[i])
		}
		return sum / float64(len(A))
	}

	prefixSum := make([]int, len(A)+1)

	for i := 1; i <= len(A); i++ {
		prefixSum[i] = prefixSum[i-1] + A[i-1]
	}
	fmt.Println(prefixSum)

	dp := make([][]float64, len(A)+1)
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]float64, K+1)
	}
	for i := 1; i <= len(A); i++ {
		dp[i][0] = float64(prefixSum[i]) / float64(i)
	}

	for k := 1; k <= K; k++ {
		for i := 1; i <= len(A); i++ {
			for j := i + 1; j <= len(A); j++ {
				prefixin := float64(prefixSum[j-1]-prefixSum[i-1]) / float64(j-i)
				dp[i][k] = math.Max(dp[i][k-1], prefixin+dp[i-1][k-1])
			}
		}
	}
	return dp[len(A)][K]
}

func largestSumOfAveragesspaceOn(A []int, K int) float64 {
	n := len(A)

	//use prefix array to calculate sum(i, j)
	P := make([]int, n+1)
	for i := 0; i < n; i++ {
		P[i+1] = A[i] + P[i]
	}
	dp := make([]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = float64(P[n]-P[i]) / float64(n-i)
	}
	for k := 1; k < K; k++ {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				dp[i] = math.Max(dp[i], dp[j]+float64(P[j]-P[i])/float64(j-i))
			}
		}
	}
	return dp[0]
}
