package DynamicProgramming

func numberOfSubarrays(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	pre := make([]int, len(nums)+1)
	cnt := make([]int, len(nums)+1)
	cnt[0] = 1
	for i := 1; i <= len(nums); i++ {
		// nums[i-1] & 1目的是判断这个是是否为基数
		pre[i] = pre[i-1] + (nums[i-1] & 1)
		if pre[i] >= k {
			res += cnt[pre[i]-k]
		}
		cnt[pre[i]]++
	}
	return res
}

func numberOfSubarraysByMap(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	preSumMap := make(map[int]int)
	preSumMap[0] = 1
	res := 0
	pre := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		// nums[i-1] & 1目的是判断这个是是否为基数
		pre[i] = pre[i-1] + (nums[i-1] & 1)
		if pre[i] >= k {
			if val, ok := preSumMap[pre[i]-k]; ok {
				res += val
			}
		}
		if val, ok := preSumMap[pre[i]]; ok {
			preSumMap[pre[i]] = val + 1
		} else {
			preSumMap[pre[i]] = 1
		}
	}
	return res
}
