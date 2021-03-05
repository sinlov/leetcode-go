### 解题思路

基本思路为 DFS 四个方向，加标记数组搜索

- 定义 四个方向的数组
- 定义一个访问标记状态的数组
- 任意一个起点开始，遍历整个 board
  - 先标记当前为状态访问过
    - 满足搜索条件，则
      - 向 4 个方向分别 DFS 搜索
    - 不满足，则重置访问状态
- 直到所有的单词字母都找到了就输出 true，否则输出 false

### 代码

```go
var direction = [][]int{ // direction
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board)) // mark visite board
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool { // in board condition
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func searchWord(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 { // not search
		return board[x][y] == word[index]
	}
	if board[x][y] == word[index] {
		visited[x][y] = true     // mark visited first
		for i := 0; i < 4; i++ { // four direction
			dx := x + direction[i][0]
			dy := y + direction[i][1]
			if isInBoard(board, dx, dy) && !visited[dx][dy] && searchWord(board, visited, word, index+1, dx, dy) {
				return true // four direction seach
			}
		}
		visited[x][y] = false // not search remark visited false
	}
	return false
}

```
