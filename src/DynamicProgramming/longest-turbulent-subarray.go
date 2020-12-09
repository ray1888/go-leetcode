package DynamicProgramming

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = 1
		dp[i][1] = 1
	}
	maxValue := 1
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			dp[i][0] = dp[i-1][1] + 1
		} else if arr[i] < arr[i-1] {
			dp[i][1] = dp[i-1][0] + 1
		}
		maxValue = max(maxValue, max(dp[i][0], dp[i][1]))
	}
	return maxValue
}
