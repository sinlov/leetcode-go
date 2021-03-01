package _00073_set_matrix_zeroes

func setZeroes(matrix [][]int) {
	row := len(matrix)    // row
	col := len(matrix[0]) // column

	firstRow := false
	firstCol := false

	for i := 0; i < row; i++ { // mark raw
		if matrix[i][0] == 0 {
			firstCol = true
			break
		}
	}
	for i := 0; i < col; i++ { // mark column
		if matrix[0][i] == 0 {
			firstRow = true
			break
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// final change
	if firstRow { // change first row
		for i := 0; i < col; i++ {
			matrix[0][i] = 0
		}
	}
	if firstCol { // change first column
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}
