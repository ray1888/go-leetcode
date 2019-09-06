package binary_search

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0]) - 1
	i := 0
	j := col
	for i < row && j >= 0 {
		if matrix[i][j] > target {
			j -= 1
		} else if matrix[i][j] < target {
			i += 1
		} else {
			return true
		}
	}
	return false
}
