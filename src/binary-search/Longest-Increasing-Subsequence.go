package binary_search

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _BinarySearchInsPos(d []int, end int, target int) int {
	start := 0
	end = end - 1
	for start <= end {
		mid := start + (end-start)/2
		if d[mid] > target {
			end = mid - 1
		} else if d[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return start
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	d := make([]int, len(nums))
	//length := len(nums)
	lengthLIS := 0
	for _, val := range nums {
		index := _BinarySearchInsPos(d, lengthLIS, val)
		d[index] = val
		if index == lengthLIS {
			lengthLIS++
		}
	}
	return lengthLIS
}

func lengthOfLISDP(nums []int) int {
	dp := make([]int, len(nums))

	maxValue := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxValue = max(maxValue, dp[i])
	}
	return maxValue
}
