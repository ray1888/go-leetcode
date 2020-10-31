package numbers

func subarraySum(nums []int, k int) int {
	countMap := map[int]int{0: 1}
	var sum, cnt int
	for _, number := range nums {
		sum += number
		if val, ok := countMap[sum-k]; ok {
			cnt += val
		}
		var sumCnt int
		if val, ok := countMap[sum]; ok {
			sumCnt = val
		}
		countMap[sum] = sumCnt + 1
	}
	return cnt
}
