package numbers

func merge(A []int, m int, B []int, n int) {
	pa := m - 1
	pb := n - 1
	tail := m + n - 1
	cur := 0
	for pa >= 0 || pb >= 0 {
		if pa == -1 {
			cur = B[pb]
			pb--
		} else if pb == -1 {
			cur = A[pa]
			pa--
		} else if A[pa] > B[pb] {
			cur = A[pa]
			pa--
		} else {
			cur = B[pb]
			pb--
		}
		A[tail] = cur
		tail--
	}
	return
}
