package _00054_spiral_matrix

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 { // special case
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
