package DynamicProgramming

func MinCost(nums [][]int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, 3)
	}
	d[0][0] = nums[0][0]
	d[0][1] = nums[0][1]
	d[0][2] = nums[0][2]

	for i := 1; i < n; i++ {
		d[i][0] = min(d[i-1][1], d[i-1][2]) + nums[i][0]
		d[i][1] = min(d[i-1][0], d[i-1][2]) + nums[i][1]
		d[i][2] = min(d[i-1][1], d[i-1][1]) + nums[i][2]
	}

	return minmin(d[n-1][0], d[n-1][1], d[n-1][2])
}
