### 解题思路

- 设置 三种缓存， row, column, sudoku 分别为 行 列 9宫格 缓存
- 遍历时，不断添加到 对应的缓存并且将占位的值置为 true
  - 每次遍历，查询是否将缓存对应的值，是否已经为 true，如果为 true 那么就不满足 sudoku 的要求（重复）
  - 完全遍历，那么就说明当前还是一个满足 sudoku 要求的 board

### 代码

```go
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
```