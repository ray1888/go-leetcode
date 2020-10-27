package numbers

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	for i := 0; i < n; i++ {
		if sum[n]-sum[i+1] == sum[i] {
			return i
		}
	}
	return -1
}
