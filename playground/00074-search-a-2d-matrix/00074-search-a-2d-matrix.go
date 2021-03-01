package _00074_search_a_2d_matrix

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
