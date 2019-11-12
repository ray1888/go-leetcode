package heap

func topKFrequent(nums []int, k int) []int {
	if len(nums) <= 0 || k <= 0 {
		return []int{}
	}
	mmap := make(map[int]int)
	for _, val := range nums {
		if _, ok := mmap[val]; !ok {
			mmap[val] = 1
		} else {
			mmap[val] += 1
		}
	}

	buckets := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		buckets[i] = make([]int, 0)
	}
	for key, val := range mmap {
		buckets[val] = append(buckets[val], key)
	}

	result := make([]int, 0, k)
	for i := len(nums); i > 0; i-- {
		for l := 0; l < len(buckets[i]); l++ {
			if len(result) < k {
				result = append(result, buckets[i][l])
			} else {
				break
			}
		}

	}
	return result

}
