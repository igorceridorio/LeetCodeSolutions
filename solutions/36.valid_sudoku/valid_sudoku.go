package validsudoku

func isValidSudoku(board [][]byte) bool {

	// Rows validation
	for i := 0; i < 9; i++ {
		boardSetRow := board[i]
		if !isValidSet(boardSetRow) {
			return false
		}
	}

	// Columns validation
	for i := 0; i < 9; i++ {
		var boardSetColumn []byte
		for j := 0; j < 9; j++ {
			boardSetColumn = append(boardSetColumn, board[j][i])
		}
		if !isValidSet(boardSetColumn) {
			return false
		}
	}

	// Squares validation
	rowOffset := 0
	for i := 0; i < 3; i++ {
		colOffset := 0
		for j := 0; j < 3; j++ {
			var boardSetSquare []byte
			for k := 0 + rowOffset; k < (3 + rowOffset); k++ {
				for l := 0 + colOffset; l < (3 + colOffset); l++ {
					boardSetSquare = append(boardSetSquare, board[k][l])
				}
			}
			if !isValidSet(boardSetSquare) {
				return false
			}
			colOffset += 3
		}
		rowOffset += 3
	}

	return true
}

func isValidSet(boardSet []byte) bool {
	setMap := make(map[byte]bool)
	for _, e := range boardSet {
		currentRune := rune(e)
		if rune(currentRune) != '.' {
			if _, ok := setMap[e]; ok {
				return false
			}
			setMap[e] = true
		}
	}
	return true
}
