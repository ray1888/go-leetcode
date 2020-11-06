package numbers

// leetcode has imple this function ,this is for function signture
func isBadVersion(n int) bool {
	return false
}

func firstBadVersion(n int) int {
	low := 1
	high := n
	for low < high {
		mid := low + (high-low)/2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
