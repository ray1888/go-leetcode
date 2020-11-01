package backtracing


import (
	"strings"
)

func solve(row int, n int, result *[][]string, board [][]string, visited [][]bool) {
	if row == n {
		var newRstring []string
		for _, item := range board {
			newRstring = append(newRstring, strings.Join(item, ""))
		}
		*result = append(*result, newRstring)
		return
	}
	for col := 0; col < n; col++ {
		// visited[0] is row base conflict, visited[1] is positive diff conflict, visited[1] is positive negetive conflict,
		if !visited[0][col] && !visited[1][row-col+n] && !visited[2][row+col] {
			board[row][col] = "Q"
			visited[0][col] = true
			visited[1][row-col+n] = true
			visited[2][row+col] = true
			solve(row+1, n, result, board, visited)
			visited[0][col] = false
			visited[1][row-col+n] = false
			visited[2][row+col] = false
			board[row][col] = "."
		}
	}
}

func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}

	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	visited := make([][]bool, 3)

	for i := 0; i < 3; i++ {
		visited[i] = make([]bool, 2*n)
	}

	result := make([][]string, 0)
	solve(0, n, &result, board, visited)
	return result
}
