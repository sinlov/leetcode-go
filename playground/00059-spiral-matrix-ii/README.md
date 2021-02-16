### 解题思路

- 排除特殊情况 n == 1 , 直接输出
- 矩阵 高 n 宽 n, 矩阵每次旋转边界坐标 `right bottom left top` 为 `n-1, n-1, 0, 0`
- 初始化输出 res 矩阵默位置，
- 循环 index < sum (总共坐标大小 sum = n * n)
  - 方向 → 循环条件 i := left; i <= right; i++ ，每次 index++ ，循环完成后收缩 top++
  - 方向 ￬ 循环条件 i := top; i <= bottom; i++ ，每次 index++ ，循环完成后收缩 right--
  - 方向 ← 循环条件 i := right; i >= left; i-- ，每次 index++ ，循环完成后收缩 bottom--
  - 方向 ￪ 循环条件 i := bottom; i >= top; i-- ，每次 index++ ，循环完成后收缩 left++

### 代码

```go
func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	res := [][]int{}
	for i := 0; i < n; i++ { // make default matrix
		res = append(res, make([]int, n))
	}
	sum := n * n // full sum number
	right, bottom, left, top := n-1, n-1, 0, 0
	index := 1 // start number
	for index <= sum {
		for i := left; i <= right; i++ { // walk →
			res[top][i] = index
			index++
		}
		top++                            // top shrink

		for i := top; i <= bottom; i++ { // walk ￬
			res[i][right] = index
			index++
		}
		right--                          // right shrink

		for i := right; i >= left; i-- { // walk ←
			res[bottom][i] = index
			index++
		}
		bottom--                         // bottom shrink

		for i := bottom; i >= top; i-- { // walk ￪
			res[i][left] = index
			index++
		}
		left++ // left shrink
	}
	return res
}
```
