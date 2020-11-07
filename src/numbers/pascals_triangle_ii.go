package numbers

func getRow(rowIndex int) []int {
	tri := make([][]int, 0)
	for i := 0; i <= rowIndex; i++ {
		r := make([]int, i+1)
		r[0] = 1
		r[i] = 1
		for j := 1; j < i; j++ {
			r[j] = tri[i-1][j] + tri[i-1][j-1]
		}
		tri = append(tri, r)
	}
	return tri[rowIndex]
}
