package backpack

import "go-leetcode/src/utils"

func backPackII(m int, A []int, V []int) int {
	dp := make([][]int, len(A)+1)
	maxValue := 0
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= len(A); i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= A[i-1] {
				dp[i][j] = utils.Max(dp[i-1][j], dp[i-1][j-A[i-1]]+V[i-1])
			}
			maxValue = utils.Max(maxValue, dp[i][j])
		}
	}
	return maxValue
}

/*
	因为如果只要一维数组的情况下，直接使用顺序遍历的话，会把同一个i之前的状态进行计算导致存储状态的失效，因此，只能使用反向减的策略来保证大容积的值会先计算
	从而保证memorize的数组可以有序。
*/
func backPackIIOnSpace(m int, A []int, V []int) int {
	dp := make([]int, m+1)

	for i := 1; i <= len(A); i++ {
		for j := m; j >= A[i-1]; j-- {
			dp[j] = utils.Max(dp[j], dp[j-A[i-1]]+V[i-1])
		}
	}
	return dp[m]
}
