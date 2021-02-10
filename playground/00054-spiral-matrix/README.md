### 解题思路

- 排除特殊情况 len(matrix) ==0 , len(matrix[0]) == 0
- 矩阵 高 m 宽 n, 矩阵每次旋转边界坐标 left right top bottom 为 0, n-1 m-1, 0
- 计数器 cnt 总共坐标大小 sum = m * n
- 循环 cnt < sum
  - 方向 → i, j := top, left ； 边界条件 j <= right && cnt < sum ，每次 j++ 并且 cnt++
  - 方向 ￬ i, j = top+1, right ；边界条件 i <= bottom && cnt < sum ，每次 i++ 并且 cnt++
  - 方向 ← i, j = bottom, right-1 ；边界条件 j >= left && cnt < sum ，每次 j-- 并且 cnt++
  - 方向 ￪ i, j = bottom, right-1 ；边界条件 i > top && cnt < sum ，每次 i-- 并且 cnt++
  - 方向再次换到 → 时，需要收缩一圈 重新 设置转边界坐标
    - left, right, bottom, top = left+1, right-1, bottom-1, top+1

### 代码

```go
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 { //special case
		return []int{}
	}
	n := len(matrix[0])
	if n == 0 { //special case
		return []int{}
	}

	left, right, bottom, top := 0, n-1, m-1, 0
	cnt, sum := 0, m*n

	res := []int{}
	for cnt < sum {
		i, j := top, left
		for j <= right && cnt < sum { // →
			res = append(res, matrix[i][j])
			cnt++
			j++
		}
		i, j = top+1, right
		for i <= bottom && cnt < sum { // ￬
			res = append(res, matrix[i][j])
			cnt++
			i++
		}
		i, j = bottom, right-1
		for j >= left && cnt < sum { // ←
			res = append(res, matrix[i][j])
			cnt++
			j--
		}
		i, j = bottom-1, left
		for i > top && cnt < sum { // ￪
			res = append(res, matrix[i][j])
			cnt++
			i--
		}
		// to inner
		left, right, bottom, top = left+1, right-1, bottom-1, top+1
	}
	return res
}
```
