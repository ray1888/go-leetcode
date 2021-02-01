package DynamicProgramming

// 此题题解

// 当时用过矩阵进行推演之后可以发现，
// 矩阵的推演信息并不能一直填写满整个矩阵（因为逻辑上面递推是相同的开头才能进行计算，因此我们需要多用一个变量去存放这个问题的最大值的解）

func findLength(A []int, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}

	la := len(A)
	lb := len(B)

	dp := make([][]int, la+1)

	for i := 0; i <= la; i++ {
		dp[i] = make([]int, lb+1)
	}

	maxValue := 0

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			maxValue = max(maxValue, dp[i][j])
		}
	}
	return maxValue
}
