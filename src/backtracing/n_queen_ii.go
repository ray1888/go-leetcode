package backtracing

func solveQueen2(row int, n int, result *int, visited [][]bool) {
	if row == n {
		(*result)++
		return
	}
	for col := 0; col < n; col++ {
		// visited[0] is row base conflict, visited[1] is positive diff conflict, visited[1] is positive negetive conflict,
		if !visited[0][col] && !visited[1][row-col+n] && !visited[2][row+col] {
			visited[0][col] = true
			visited[1][row-col+n] = true
			visited[2][row+col] = true
			solveQueen2(row+1, n, result, visited)
			visited[0][col] = false
			visited[1][row-col+n] = false
			visited[2][row+col] = false
		}
	}
}

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}

	visited := make([][]bool, 3)

	for i := 0; i < 3; i++ {
		visited[i] = make([]bool, 2*n)
	}

	result := 0
	solveQueen2(0, n, &result, visited)
	return result
}
