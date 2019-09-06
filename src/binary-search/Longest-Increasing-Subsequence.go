package binary_search

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
			lengthLIS += 1
		}
	}
	return lengthLIS
}
