package stringProblem

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	aa := A + A
	n := len(B)
	nn := len(A) * 2
	for start := 0; start <= n; start++ {
		i := start
		j := 0
		for i < nn && j < n && aa[i] == B[j] {
			i++
			j++
		}
		if j == n {
			return true
		}
	}
	return false
}
