package _00059_spiral_matrix_ii

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
