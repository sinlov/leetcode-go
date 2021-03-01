### 解题思路

- 降低复杂度，直接使用 二分搜索
- 初始化 长度 m := len(matrix[0])
- `low, hight` 初始化为 `0, len(matrix[0])*len(matrix)-1`
- 二分搜索遍历时
  - `matrix[mid/m][mid%m] == target` 找到目标
  - `matrix[mid/m][mid%m] > target` 向 low 方向找
  - `matrix[mid/m][mid%m] < target` 向 hight 方向找

### 代码

```go
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 { // special case
		return false
	}
	m := len(matrix[0])
	low, high := 0, len(matrix[0])*len(matrix)-1
	for low <= high { // binary search
		mid := low + (high-low)>>1
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] > target { // move to low
			high = mid - 1
		} else { // move to high
			low = mid + 1
		}
	}
	return false
}
```
