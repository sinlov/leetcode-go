### 解题思路

使用 `动态规划 dp`

- 累加第一列
- 累加第一行
- 其他部分按照 更小的值累加
- 最后输出 grid[m-1][n-1] 就是最小值

### 代码

```go
func minPathSum(grid [][]int) int {
	if len(grid) == 0 { // special case
		return 0
	}
	m, n := len(grid), len(grid[0])
	if n == 0 { // special case
		return 0
	}

	for i := 1; i < m; i++ { // accumulation first column
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ { // accumulation first row
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < m; i++ { // accumulation other by lessThan
		for j := 1; j < n; j++ {
			grid[i][j] += lessThan(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1] // last accumulation is minimum
}

func lessThan(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```
