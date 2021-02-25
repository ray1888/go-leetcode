package power

func isPowerOfK(n int, k int) bool {
	if n == 0 {
		return false
	}
	for n%k == 0 {
		n /= k
	}
	return n == 1
}
