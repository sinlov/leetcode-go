package _00036_valid_sudoku

func isValidSudoku(board [][]byte) bool {
	rowbuf, colbuf, sudokubuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowbuf[i] = make([]bool, 9)
		colbuf[i] = make([]bool, 9)
		sudokubuf[i] = make([]bool, 9)
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ { // range full board r = row c = column
			if board[r][c] != '.' { // number
				num := board[r][c] - '0' - byte(1)                                 // to number
				if rowbuf[r][num] || colbuf[c][num] || sudokubuf[r/3*3+c/3][num] { // has any cache is not sudoku
					return false
				}
				rowbuf[r][num] = true            // cache row
				colbuf[c][num] = true            // cache column
				sudokubuf[r/3*3+c/3][num] = true // cache sudoku
			}
		}
	}
	return true
}
