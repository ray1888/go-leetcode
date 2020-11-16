package numbers

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) == 0 || len(nums) < 2 {
		return false
	}
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		var mod int
		if k == 0 {
			mod = sum
		} else {
			mod = sum % k
		}
		if j, ok := m[mod]; ok {
			if i-j > 1 {
				return true
			}
		} else {
			m[mod] = i
		}
	}
	return false
}
