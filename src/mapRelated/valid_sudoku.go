package mapRelated

func isValidSudoku(board [][]byte) bool {
	rowSeen := make([][]bool, 9*9)
	for i := 0; i < 9; i++ {
		rowSeen[i] = make([]bool, 9)
	}
	colSeen := make([][]bool, 9*9)
	for i := 0; i < 9; i++ {
		colSeen[i] = make([]bool, 9)
	}
	boxSeen := make([][]bool, 9*9)
	for i := 0; i < 9; i++ {
		boxSeen[i] = make([]bool, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if uint8(board[i][j]) == '.' {
				continue
			}
			x := uint8(board[i][j]) - '1'
			k := (i/3)*3 + j/3
			if rowSeen[i][x] || colSeen[j][x] || boxSeen[k][x] {
				return false
			}
			rowSeen[i][x] = true
			colSeen[j][x] = true
			boxSeen[k][x] = true
		}
	}
	return true
}
