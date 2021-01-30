package numbers

func searchMatrixii(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		if target == 0 {
			return true
		}
		return false
	}
	row := len(matrix)
	col := len(matrix[0]) - 1
	i := 0
	j := col
	for i < row && j >= 0 {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}
	return false
}
