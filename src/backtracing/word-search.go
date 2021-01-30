package backtracing

func exist(board [][]byte, word string) bool {
	if len(board) == 0 && len(word) == 0 {
		return true
	}
	if len(board) == 0 {
		return false
	}
	row := len(board)
	col := len(board[0])

	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if wordExistDfs(board, visited, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func wordExistDfs(board [][]byte, visited [][]bool, start, end int, word string, idx int) bool {
	if idx >= len(word) {
		return true
	}
	if start < 0 || start >= len(board) || end >= len(board[0]) || end < 0 || visited[start][end] || board[start][end] != word[idx] {
		return false
	}
	visited[start][end] = true
	existed := (wordExistDfs(board, visited, start+1, end, word, idx+1) ||
		wordExistDfs(board, visited, start, end-1, word, idx+1) ||
		wordExistDfs(board, visited, start-1, end, word, idx+1) ||
		wordExistDfs(board, visited, start, end+1, word, idx+1))
	visited[start][end] = false
	return existed
}
