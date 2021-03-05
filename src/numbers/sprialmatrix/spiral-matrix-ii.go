package sprialmatrix

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	nums := 1
	tar := n * n
	l := 0
	t := 0
	r := n - 1
	b := n - 1
	for nums <= tar {
		for i := l; i <= r; i++ {
			matrix[t][i] = nums
			nums++
		}
		t++
		for i := t; i <= b; i++ {
			matrix[i][r] = nums
			nums++
		}
		r--
		for i := r; i >= l; i-- {
			matrix[b][i] = nums
			nums++
		}
		b--
		for i := b; i >= t; i-- {
			matrix[i][l] = nums
			nums++
		}
		l++
	}
	return matrix
}
