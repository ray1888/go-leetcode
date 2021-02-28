package numbers

func isMonotonic(A []int) bool {

	isAsc := true
	isDesc := true
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			isAsc = false
			break
		}
	}

	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			isDesc = false
			break
		}
	}

	return isDesc || isAsc
}

func isMonotonicOn(A []int) bool {
	inc, dec := true, true
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			inc = false
		}
		if A[i] < A[i+1] {
			dec = false
		}
	}
	return inc || dec
}
